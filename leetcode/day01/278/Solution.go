package main

import "fmt"

func main() {
	res := firstBadVersion(5)
	fmt.Println(res)
}
func isBadVersion(version int) bool {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	return nums[version] >= 1
}

func firstBadVersion(n int) int {
	l, r := 0, 9
	min := r
	for l <= r {
		mid := l + (r-l)/2
		// l - mid - r mid
		if isBadVersion(mid) && min > mid {
			min = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
		fmt.Printf("l: %d, r: %d, min: %d, mid: %d\n", l, r, min, mid)
	}
	return min
}
