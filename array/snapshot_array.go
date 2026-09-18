package main

type versionedValue struct {
	val        int
	snapshotID int
}

// SnapshotArray supports taking point-in-time snapshots of an array and
// reading back the value an index held at any past snapshot.
//
// Instead of copying the whole array on every Snap (expensive), it keeps
// per-index history: history[index] is a list of (val, snapshotID) pairs
// sorted by snapshotID, recording only the snapshots at which that index's
// value actually changed. Get then binary-searches that per-index history
// for the most recent change at or before the requested snapshot.
//
// Example:
//
//	sa := Constructor1(3)
//	sa.Set(0, 5)     // index 0 = 5 (still snapshot 0, not yet taken)
//	id := sa.Snap()  // id == 0; freezes "index 0 = 5" as of snapshot 0
//	sa.Set(0, 7)     // index 0 = 7 (belongs to snapshot 1, in progress)
//	sa.Get(0, id)    // returns 5, the value at snapshot 0
type SnapshotArray struct {
	history        [][]versionedValue
	currSnapshotID int
}

func Constructor1(length int) SnapshotArray {
	return SnapshotArray{
		history:        make([][]versionedValue, length),
		currSnapshotID: 0,
	}
}

// Set writes val to index for the current (not-yet-snapshotted) snapshot ID.
//
// If the last recorded entry for this index already belongs to the current
// snapshot ID (i.e. Set was already called on this index since the last
// Snap), it is overwritten in place rather than appending a new entry -
// only one value per index needs to survive per snapshot. Otherwise a new
// entry is appended, marking the point where this index's value changed.
func (this *SnapshotArray) Set(index int, val int) {
	lastIdx := len(this.history[index]) - 1
	if len(this.history[index]) == 0 || this.history[index][lastIdx].snapshotID < this.currSnapshotID {
		this.history[index] = append(this.history[index], versionedValue{val, this.currSnapshotID})
	} else {
		this.history[index][lastIdx] = versionedValue{val, this.currSnapshotID}
	}
}

// Snap freezes the array's current state as a new snapshot and returns its
// ID. It doesn't copy any data - the per-index history already stores what
// value each index had by the time this snapshot is taken. Subsequent Set
// calls bump currSnapshotID first (implicitly, by comparison in Set), so
// they're recorded as belonging to the next snapshot, not this one.
func (this *SnapshotArray) Snap() int {
	snapshotID := this.currSnapshotID
	this.currSnapshotID += 1
	return snapshotID
}

// Get returns the value index held at the given snap_id.
//
// history[index] is sorted by snapshotID, so this binary-searches for the
// rightmost entry whose snapshotID <= snap_id (i.e. the last value set at
// or before that snapshot was taken) - any entry with a later snapshotID
// happened after snap_id and doesn't apply yet. r ends up pointing at that
// entry; if r < 0 nothing was ever set for this index by snap_id, so the
// default zero value is returned.
func (this *SnapshotArray) Get(index int, snap_id int) int {
	l := 0
	r := len(this.history[index]) - 1
	for l <= r {
		mid := (l + r) / 2
		if this.history[index][mid].snapshotID <= snap_id {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if r < 0 {
		return 0
	}
	return this.history[index][r].val
}

// https://leetcode.com/problems/snapshot-array/description/
func snapshotArray() {
	sa := Constructor1(3)

	sa.Set(0, 5)
	sa.Set(0, 7)
	sa.Snap()
	sa.Get(0, 0)

	// less optimized code
	sa2 := Constructor2(3)

	sa2.Set(0, 5)
	sa2.Set(0, 7)
	sa2.Snap()
	sa2.Get(0, 0)
}

// less optimized code
type SnapshotArray2 struct {
	snapId int
	values []map[int]int
}

func Constructor2(length int) SnapshotArray2 {
	sa := SnapshotArray2{
		values: make([]map[int]int, length),
	}

	for i := range sa.values {
		sa.values[i] = make(map[int]int)
	}

	return sa
}

func (this *SnapshotArray2) Set(index int, val int) {
	this.values[index][this.snapId] = val
}

func (this *SnapshotArray2) Snap() int {
	this.snapId++

	return this.snapId - 1
}

func (this *SnapshotArray2) Get(index int, snap_id int) int {
	for i := snap_id; i >= 0; i-- {
		val, ok := this.values[index][i]
		if ok {
			return val
		}
	}

	return 0
}
