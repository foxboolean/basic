package main

// s1 的组合排列用 need 包含，特别注意是连续的，仅包含小写字符利用数组处理
// s2 win 滑窗处理
func checkInclusion(s1 string, s2 string) bool {
	need, window := [26]int{}, [26]int{} // 两个长度为26的数组
	for _, s := range s1 {
		need[s-'a']++
	}
	start := 0
	for i, c := range s2 {
		in := c - 'a'
		// 插入
		window[in]++
		// 需要收缩
		for start <= i && window[in] > need[in] {
			window[s2[start]-'a']--
			start++
		}

		if window == need {
			return true
		}
	}
	return false
}
