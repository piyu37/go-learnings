package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// can getHits be called with anytimestamp or timestamp would increasing from hit(ts)?
// assumption: always increasing => but try both scenario

// 2 ways we can handle:
// what is important for us? Fater Write or faster read?
// faster write => less space; use BS
// faster read => extra space; but O(1) read time
// assumption: go with less space, faster write & slightly slow read O(logn)

// variable timeperiod
// need to see how to reset increaing count no.
type HitCounter struct {
	Queue []TSCounter
}

type TSCounter struct {
	Timestamp int
	Count     int
}

func Constructor() HitCounter {
	return HitCounter{
		Queue: make([]TSCounter, 0),
	}
}

func (this *HitCounter) Hit(timestamp int) {
	if len(this.Queue) == 0 {
		this.Queue = append(this.Queue, TSCounter{timestamp, 1})
		return
	}

	last := this.Queue[len(this.Queue)-1]
	if last.Timestamp == timestamp {
		last.Count++

		this.Queue[len(this.Queue)-1] = last
		return
	}

	this.Queue = append(this.Queue, TSCounter{timestamp, last.Count + 1})
}

func (this *HitCounter) GetHits(timestamp int) int {
	if len(this.Queue) == 0 {
		return 0
	}

	last := this.Queue[len(this.Queue)-1]

	startDurationTime := timestamp - 300

	if startDurationTime <= 0 {
		return last.Count
	}

	nearestRightCount := this.getNearestRightCount(startDurationTime)
	return last.Count - nearestRightCount
}

func (this *HitCounter) getNearestRightCount(startDurationTime int) int {
	start := 0
	end := len(this.Queue) - 1

	count := 0

	for start <= end {
		mid := start + (end-start)/2

		if this.Queue[mid].Timestamp == startDurationTime {
			return this.Queue[mid].Count
		}

		if this.Queue[mid].Timestamp < startDurationTime {
			count = this.Queue[mid].Count
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return count
}

func mySoln() {
	obj := Constructor()
	obj.Hit(1)
	obj.Hit(2)
	obj.Hit(3)

	fmt.Println(obj.GetHits(4))

	obj.Hit(300)

	fmt.Println(obj.GetHits(300))
	fmt.Println(obj.GetHits(301))
}

type HitCounterUsingFixedLengthArray struct {
	Queue []TSCounter
	Len   int
}

func ConstructorUsingFixedLengthArray() HitCounterUsingFixedLengthArray {
	return HitCounterUsingFixedLengthArray{
		Queue: make([]TSCounter, 300),
		Len:   300,
	}
}

func (this *HitCounterUsingFixedLengthArray) Hit(timestamp int) {
	ts := timestamp % this.Len

	if this.Queue[ts].Timestamp == timestamp {
		this.Queue[ts].Count++
		return
	}

	this.Queue[ts] = TSCounter{Timestamp: timestamp, Count: 1}
}

func (this *HitCounterUsingFixedLengthArray) GetHits(timestamp int) int {
	count := 0
	for i := range this.Len {
		if timestamp-this.Queue[i].Timestamp < this.Len {
			count += this.Queue[i].Count
		}
	}

	return count
}

// https://leetcode.com/problems/design-hit-counter/
func hitCounter() {
	mySoln()
	fmt.Println("-----------------")
	optimizedSolutionUsingMutex()
	fmt.Println("-----------------")
	optimizedSolutionUsingAtomicPointer()
}

type HitCounterWithMutex struct {
	mu      [300]sync.RWMutex
	buckets [300]TSCounter
}

func ConstructorMutex() HitCounterWithMutex {
	return HitCounterWithMutex{}
}

func (h *HitCounterWithMutex) Hit(timestamp int) {
	idx := timestamp % 300
	h.mu[idx].Lock()
	defer h.mu[idx].Unlock()

	if h.buckets[idx].Timestamp == timestamp {
		h.buckets[idx].Count++
	} else {
		h.buckets[idx] = TSCounter{timestamp, 1}
	}
}

func (h *HitCounterWithMutex) GetHits(timestamp int) int {
	count := 0
	for i := range 300 {
		h.mu[i].RLock()
		if timestamp-h.buckets[i].Timestamp < 300 {
			count += h.buckets[i].Count
		}

		h.mu[i].RUnlock()
	}

	return count
}

func optimizedSolutionUsingMutex() {
	obj := ConstructorMutex()
	obj.Hit(1)
	obj.Hit(2)
	obj.Hit(3)

	fmt.Println(obj.GetHits(4))

	obj.Hit(300)

	fmt.Println(obj.GetHits(300))
	fmt.Println(obj.GetHits(301))
}

type HitCounterWithAtomicPointer struct {
	buckets [300]atomic.Pointer[TSCounter]
}

func Constructor2() HitCounterWithAtomicPointer {
	return HitCounterWithAtomicPointer{}
}

func (h *HitCounterWithAtomicPointer) Hit(timestamp int) {
	idx := timestamp % 300
	for {
		old := h.buckets[idx].Load()
		if old != nil && old.Timestamp == timestamp {
			newVal := &TSCounter{timestamp, old.Count + 1}
			if h.buckets[idx].CompareAndSwap(old, newVal) {
				return
			}
			continue // another goroutine updated it first, retry
		}
		newVal := &TSCounter{timestamp, 1}
		if h.buckets[idx].CompareAndSwap(old, newVal) {
			return
		}
	}
}

func (this *HitCounterWithAtomicPointer) GetHits(timestamp int) int {
	count := 0
	for i := range 300 {
		b := this.buckets[i].Load()
		if b == nil {
			continue
		}
		if timestamp-b.Timestamp < 300 {
			count += b.Count
		}
	}
	return count
}

func optimizedSolutionUsingAtomicPointer() {
	obj := Constructor2()
	obj.Hit(1)
	obj.Hit(2)
	obj.Hit(3)

	fmt.Println(obj.GetHits(4))

	obj.Hit(300)

	fmt.Println(obj.GetHits(300))
	fmt.Println(obj.GetHits(301))
}

type HitCounterWithAtomicBitPacking struct {
	buckets [300]atomic.Uint64 // top 32 bits = timestamp, bottom 32 bits = count
}

func Constructor3() HitCounterWithAtomicBitPacking {
	return HitCounterWithAtomicBitPacking{}
}

// binary:  0000...0000 0101        (64 bits total, only last 3 bits are "101")
// hex:     0000 0000  0000 0005    (16 hex digits, grouped in fours: 1 hex digit = 4 bits)

// pack squeezes two uint32 values into one uint64 by putting timestamp in the
// high 32 bits and count in the low 32 bits, so a single atomic.Uint64 can
// hold both fields and be updated with one CompareAndSwap.
//
// Example: timestamp = 5, count = 3
//
//	uint64(timestamp) = 0x00000000_00000005
//	<<32 shifts it 32 bits left, moving it into the upper half:
//	                  = 0x00000005_00000000
//	uint64(count)     = 0x00000000_00000003   (already in the lower half)
//	OR (|) combines them since the two occupy disjoint bit ranges:
//	                    0x00000005_00000000
//	                  | 0x00000000_00000003
//	                  = 0x00000005_00000003
//
// ts = 5, count = 3
func pack(timestamp, count uint32) uint64 {
	return uint64(timestamp)<<32 | uint64(count)
}

// unpack reverses pack: shift right to pull the timestamp back out of the
// high bits, and truncate to uint32 to read just the count from the low bits.
//
// Example: v = 0x00000005_00000003 (timestamp=5, count=3)
//
//	v >> 32 shifts the high 32 bits down into the low 32 bits, discarding
//	the old low bits off the end:
//	          0x00000005_00000003 >> 32 = 0x00000000_00000005
//	uint32(...) keeps only the low 32 bits of that     = 5  (timestamp)
//
//	uint32(v) truncates v directly to its low 32 bits, dropping the high
//	bits entirely (no shift needed since count already lives at the bottom):
//	          uint32(0x00000005_00000003) = 0x00000003 = 3  (count)
func unpack(v uint64) (timestamp, count uint32) {
	return uint32(v >> 32), uint32(v)
}
func (h *HitCounterWithAtomicBitPacking) Hit(timestamp int) {
	idx := timestamp % 300
	ts := uint32(timestamp)
	for {
		old := h.buckets[idx].Load()
		oldTs, oldCount := unpack(old)
		var newVal uint64
		if old != 0 && oldTs == ts {
			newVal = pack(ts, oldCount+1)
		} else {
			newVal = pack(ts, 1)
		}
		if h.buckets[idx].CompareAndSwap(old, newVal) {
			return
		}
	}
}

func (this *HitCounterWithAtomicBitPacking) GetHits(timestamp int) int {
	totalCount := 0
	for i := range 300 {
		b := this.buckets[i].Load()
		ts, count := unpack(b)
		if timestamp-int(ts) < 300 {
			totalCount += int(count)
		}
	}

	return totalCount
}

func optimizedSolutionUsingAtomicBitPacking() {
	obj := Constructor3()
	obj.Hit(1)
	obj.Hit(2)
	obj.Hit(3)

	fmt.Println(obj.GetHits(4))

	obj.Hit(300)

	fmt.Println(obj.GetHits(300))
	fmt.Println(obj.GetHits(301))
}
