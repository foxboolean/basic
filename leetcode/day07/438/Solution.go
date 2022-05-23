package _38

// 仅包含小写字母使用数组存储
// 包含异位词，滑动窗口
func findAnagrams(s string, p string) []int {
	need, window := [26]int{}, [26]int{}
	for _, c := range p {
		need[c-'a']++
	}
	res := make([]int, 0)
	start := 0
	for i, c := range s {
		// 右侧扩充
		window[c-'a']++
		// 左侧收窄
		for start <= i && window[s[start]-'a'] > need[s[start]-'a'] {
			window[s[start]-'a']--
			start++
		}
		if need == window {
			res = append(res, start)
		}
	}
	return res
}
