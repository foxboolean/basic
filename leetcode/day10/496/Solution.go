package _96

//单调栈求 nums2 的 next greater
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var next []int
	m := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(next) > 0 && next[len(next)-1] < num {
			next = next[0 : len(next)-1]
		}
		if len(next) > 0 {
			m[num] = next[len(next)-1]
		} else {
			m[num] = -1
		}
		next = append(next, num)
	}

	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = m[v]
	}
	return res
}
