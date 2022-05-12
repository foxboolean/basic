package main

import "fmt"

func main() {
	//res := removeDuplicates([]int{1, 1, 2})
	res := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Printf("res: %d", res)
}

// 有序数组，去重
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		} else {
			fast++
		}
	}
	return slow + 1
}
