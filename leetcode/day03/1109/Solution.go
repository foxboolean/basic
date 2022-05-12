package _109

// 差分数组，还原数组
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	// first, last, n
	for _, booking := range bookings {
		diff[booking[0]-1] += booking[2]
		if booking[1] < n {
			diff[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff
}
