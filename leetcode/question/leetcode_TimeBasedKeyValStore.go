// 981. Time Based Key-Value Store
// 透過有序timestamp, 可以用二分搜尋

type entry struct {
	time  int
	value string
}

type TimeMap struct {
	store map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]entry)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.store[key] = append(tm.store[key], entry{timestamp, value})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	entries := tm.store[key]
	if len(entries) == 0 {
		return ""
	}
	// binary search
	left, right := 0, len(entries)-1
	res := ""
	for left <= right {
		mid := (left + right) / 2
		if entries[mid].time <= timestamp {
			res = entries[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}