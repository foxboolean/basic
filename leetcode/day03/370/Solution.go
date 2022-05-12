package main

// getModifiedArray updates [start, end, inc]
func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length)
	for _, update := range updates {
		diff[update[0]] += update[2]
		if update[1] < length-1 {
			diff[update[1]+1] -= update[2]
		}
	}
	for i := 1; i < length; i++ {
		diff[i] = diff[i-1] + diff[i]
	}
	return diff
}
