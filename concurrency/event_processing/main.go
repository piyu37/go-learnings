package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// ---------------------------------------------------------------------------
// 1. Input: one row from the job_run history table.
// ---------------------------------------------------------------------------

// Record is exactly what the question describes. FinishedAt is a pointer
// so it can be nil, which is how we represent SQL NULL ("the worker died").
type Record struct {
	JobRunID       string
	JobType        string
	Attempt        int
	StartedAt      time.Time
	FinishedAt     *time.Time // nil = worker died, never reported finishing
	LeaseExpiresAt time.Time
}

// ---------------------------------------------------------------------------
// 2. Turn each record into a time interval [Start, End).
// ---------------------------------------------------------------------------

// Attempt is the cleaned-up version: just a start and an end.
// The interval is half-open: Start is included, End is NOT.
// So [10:00, 10:05) and [10:05, 10:06) do not overlap.
type Attempt struct {
	JobRunID string
	JobType  string
	Start    time.Time
	End      time.Time
}

// toAttempt applies the NULL rule:
// if the worker never reported finishing, assume it stopped at lease expiry.
func toAttempt(r Record) Attempt {
	end := r.LeaseExpiresAt
	if r.FinishedAt != nil {
		end = *r.FinishedAt
	}
	return Attempt{
		JobRunID: r.JobRunID,
		JobType:  r.JobType,
		Start:    r.StartedAt,
		End:      end,
	}
}

// ---------------------------------------------------------------------------
// 3. Part (a): which job_run_ids had overlapping attempts?
// ---------------------------------------------------------------------------

func findOverlaps(attempts []Attempt) []string {
	// Sort by (JobRunID, Start) so each job's attempts sit together, in order.
	sort.Slice(attempts, func(i, j int) bool {
		if attempts[i].JobRunID != attempts[j].JobRunID {
			return attempts[i].JobRunID < attempts[j].JobRunID
		}
		return attempts[i].Start.Before(attempts[j].Start)
	})

	var result []string
	var maxEnd time.Time // latest end seen so far for the current job

	for i, a := range attempts {
		newJob := i == 0 || a.JobRunID != attempts[i-1].JobRunID

		if newJob {
			// First attempt of a job: nothing to compare against yet.
			maxEnd = a.End
			continue
		}

		// Overlap if this attempt starts before an earlier one has ended.
		// Strictly "Before": starting exactly at maxEnd is a clean handoff.
		if a.Start.Before(maxEnd) {
			// Report the job once, even if it has several overlapping attempts.
			if len(result) == 0 || result[len(result)-1] != a.JobRunID {
				result = append(result, a.JobRunID)
			}
		}

		if a.End.After(maxEnd) {
			maxEnd = a.End
		}
	}
	return result
}

// ---------------------------------------------------------------------------
// 4. Part (b): peak concurrency per job_type (the sweep line).
// ---------------------------------------------------------------------------

// Event is one endpoint of an attempt: +1 when it starts, -1 when it ends.
type Event struct {
	T       time.Time // when this event happens
	Delta   int       // +1 start, -1 end
	JobType string
}

type Peak struct {
	Count int
	At    time.Time
}

func peakConcurrency(attempts []Attempt) map[string]Peak {
	// Each attempt becomes two events, so 2n events in total.
	events := make([]Event, 0, 2*len(attempts))
	for _, a := range attempts {
		events = append(events, Event{T: a.Start, Delta: +1, JobType: a.JobType})
		events = append(events, Event{T: a.End, Delta: -1, JobType: a.JobType})
	}

	// Sort by time. At the same time, ends (-1) come before starts (+1),
	// because a job ending at 10:05 and one starting at 10:05 never ran together.
	sort.Slice(events, func(i, j int) bool {
		if !events[i].T.Equal(events[j].T) {
			return events[i].T.Before(events[j].T)
		}
		return events[i].Delta < events[j].Delta
	})

	running := map[string]int{} // how many attempts are running right now, per type
	peaks := map[string]Peak{}  // best seen so far, per type

	for _, e := range events {
		running[e.JobType] += e.Delta
		now := running[e.JobType]

		// Use ">" so we keep the FIRST moment the peak was reached.
		if now > peaks[e.JobType].Count {
			peaks[e.JobType] = Peak{Count: now, At: e.T}
		}
	}
	return peaks
}

// ---------------------------------------------------------------------------
// 5. Small hand-made test dataset.
// ---------------------------------------------------------------------------

// The test data is written as plain "HH:MM" strings so it is easy to read.
// "NULL" in the finished column means the worker died.
//
//	job    type      attempt  started  finished  lease
var rows = [][]string{
	// J1: attempt 1 outran its lease (10:05) and finished at 10:08,
	// but attempt 2 started at 10:05. They overlap from 10:05 to 10:07.
	{"J1", "email", "1", "10:00", "10:08", "10:05"},
	{"J1", "email", "2", "10:05", "10:07", "10:10"},

	// J2: worker died, so attempt 1 ends at its lease (10:05).
	// Attempt 2 starts exactly at 10:05: clean handoff, NOT an overlap.
	{"J2", "email", "1", "10:00", "NULL", "10:05"},
	{"J2", "email", "2", "10:05", "10:06", "10:10"},

	// J4: single attempt, nothing to overlap with.
	{"J4", "email", "1", "10:10", "10:15", "10:20"},

	// J5: worker died, assumed running until its lease (10:12).
	// Attempt 2 started at 10:09, so they overlap.
	{"J5", "email", "1", "10:02", "NULL", "10:12"},
	{"J5", "email", "2", "10:09", "10:11", "10:15"},

	// R1: three attempts; attempt 1 is long and overlaps both others.
	{"R1", "resize", "1", "11:00", "11:20", "11:10"},
	{"R1", "resize", "2", "11:10", "11:15", "11:20"},
	{"R1", "resize", "3", "11:12", "11:18", "11:22"},

	// R2, R3: single attempts. R3 starts exactly when R1 attempt 1 ends.
	{"R2", "resize", "1", "11:05", "11:09", "11:15"},
	{"R3", "resize", "1", "11:20", "11:25", "11:30"},
}

// parseHM turns "10:05" into a time.Time.
func parseHM(s string) time.Time {
	t, err := time.Parse("15:04", s)
	if err != nil {
		panic(err)
	}
	return t
}

// testData converts the string rows above into Records.
func testData() []Record {
	var records []Record
	for _, r := range rows {
		attempt, _ := strconv.Atoi(r[2])

		var finished *time.Time // stays nil for "NULL"
		if r[4] != "NULL" {
			t := parseHM(r[4])
			finished = &t
		}

		records = append(records, Record{
			JobRunID:       r[0],
			JobType:        r[1],
			Attempt:        attempt,
			StartedAt:      parseHM(r[3]),
			FinishedAt:     finished,
			LeaseExpiresAt: parseHM(r[5]),
		})
	}
	return records
}

// As we wil get big data of events, solve this via concurrency

// ===========================================================================
// PROBLEM: Find overlapping job attempts and peak concurrency from lease records
// ===========================================================================
//
// BACKGROUND
//
//	Background jobs run on worker machines. When a worker picks up a job it
//	gets a LEASE (a deadline). If the lease expires before the worker reports
//	back, the scheduler assumes it crashed and starts the job again on another
//	worker. If the first worker was only slow (not dead), the same job now
//	runs TWICE at the same time. Each try is called an "attempt".
//
// INPUT (one day ~ 50 million rows)
//
//	job_run_id, job_type, attempt, started_at, finished_at, lease_expires_at
//	finished_at is NULL when the worker died and never reported finishing.
//
// OUTPUT
//
//	(a) Every job_run_id whose attempts overlapped in time
//	    (the job really ran twice at once).
//	(b) For each job_type, the peak number of attempts running at the same
//	    moment, and the minute it happened.
//
// CONSTRAINT
//
//	Must be O(n log n). Comparing every pair is O(n^2), too slow for 50M rows.
//
// KEY DECISIONS
//
//  1. Interval of an attempt = [started_at, end), where
//     end = finished_at, or lease_expires_at if finished_at is NULL.
//     A dead worker might have kept running past its lease, so results are
//     a LOWER BOUND: real overlap is at least what we report.
//
//  2. Half-open interval [start, end): an attempt ending at 10:05 and one
//     starting at 10:05 did NOT run together (clean handoff).
//
//  3. Clock skew: timestamps come from different machines whose clocks
//     differ slightly. Tiny overlaps (smaller than the possible clock error)
//     may be fake, and tiny real overlaps may be hidden. Treat them as
//     "uncertain", not definite.
//
// APPROACH
//
//	(a) Sort attempts by (job_run_id, started_at). Walk once, tracking the
//	    latest end time seen for the current job. If the next attempt starts
//	    before that end -> overlap.
//
//	(b) Sweep line: each attempt becomes two events, +1 at start and -1 at
//	    end. Sort events by time; on ties put END before START. Walk once,
//	    keeping a running count per job_type and recording each type's max
//	    and when it occurred (truncated to the minute).
//
//	Both parts: sort O(n log n) + one linear pass O(n).
//
// ===========================================================================
// https://prachub.com/interview-guide/arctic-wolf-software-engineer-interview-questions-guide-2026#:~:text=Find%20overlapping%20job%20attempts%20and%20peak%20concurrency%20from%20lease%20records
func main() {
	records := testData()

	attempts := make([]Attempt, len(records))
	for i, r := range records {
		attempts[i] = toAttempt(r)
	}

	// Part (a) sorts its input, so give it its own copy.
	forOverlaps := make([]Attempt, len(attempts))
	copy(forOverlaps, attempts)

	fmt.Println("(a) Jobs whose attempts overlapped:")
	for _, id := range findOverlaps(forOverlaps) {
		fmt.Println("   ", id)
	}

	fmt.Println()
	fmt.Println("(b) Peak concurrency per job_type:")
	peaks := peakConcurrency(attempts)

	// Print types in a fixed order (Go map order is random).
	types := make([]string, 0, len(peaks))
	for t := range peaks {
		types = append(types, t)
	}
	sort.Strings(types)

	for _, t := range types {
		p := peaks[t]
		fmt.Printf("    %-7s peak=%d at minute %s\n",
			t, p.Count, p.At.Truncate(time.Minute).Format("15:04"))
	}
}
