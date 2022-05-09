package main

import "fmt"

func main() {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	//target := 9
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	res := search(nums, target)
	fmt.Println(res)
}

// 升序
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1 // 重点
		} else if nums[mid] > target {
			r = mid - 1 // 重点
		}
	}
	return -1
}
