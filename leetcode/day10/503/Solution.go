package main

func main() {
	println(nextGreaterElements([]int{1, 2, 3}))
}

// 循环数组 -> 下标扩充 「数组扩充占用内存」
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	var s []int
	res := make([]int, n*2)
	for i := n*2 - 1; i >= 0; i-- {
		num := nums[i%n]
		for len(s) > 0 && num >= s[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) > 0 {
			res[i] = s[len(s)-1]
		} else {
			res[i] = -1
		}
		s = append(s, num)
	}

	return res[0:n]
}
