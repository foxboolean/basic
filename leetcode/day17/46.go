package day17

var ans [][]int

func permute(nums []int) [][]int {
	ans = [][]int{}
	backtrack(nums, []int{})
	return ans
}

func backtrack(nums []int, s []int) {
	sz := len(nums)
	if sz == 0 {
		p := make([]int, len(s))
		copy(p, s) // 指向同一个区域避免值被修改
		ans = append(ans, p)
	}

	for i := 0; i < sz; i++ {
		curr := nums[i]
		s = append(s, curr)
		// 移除已经使用的元素
		nums = append(nums[:i], nums[i+1:]...)
		backtrack(nums, s)
		// 还原元素
		nums = append(nums[:i], append([]int{curr}, nums[i:]...)...)
		s = s[:len(s)-1]
	}
}
