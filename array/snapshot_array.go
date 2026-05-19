package main

type versionedValue struct {
	val        int
	snapshotID int
}

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

func (this *SnapshotArray) Set(index int, val int) {
	lastIdx := len(this.history[index]) - 1
	if len(this.history[index]) == 0 || this.history[index][lastIdx].snapshotID < this.currSnapshotID {
		this.history[index] = append(this.history[index], versionedValue{val, this.currSnapshotID})
	} else {
		this.history[index][lastIdx] = versionedValue{val, this.currSnapshotID}
	}
}

func (this *SnapshotArray) Snap() int {
	snapshotID := this.currSnapshotID
	this.currSnapshotID += 1
	return snapshotID
}

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
}
