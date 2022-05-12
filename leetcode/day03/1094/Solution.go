package _094

// 构造差分数组，还原原始数组
func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1000)
	// num, from, to
	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		if trip[2] < 1000 {
			diff[trip[2]] -= trip[0]
		}
	}

	for i := 0; i < len(diff); i++ {
		if i != 0 {
			diff[i] = diff[i-1] + diff[i]
		}
		if diff[i] > capacity {
			return false
		}
	}
	return true
}
