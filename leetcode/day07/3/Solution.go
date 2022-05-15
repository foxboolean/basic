package main

func main() {
	substring := lengthOfLongestSubstring("bbbbb")
	println(substring)
}

// 不含重复字符的最长子串 -> 滑动窗口
func lengthOfLongestSubstring(s string) int {
	window := make(map[string]int, len(s))
	res, l, r := 0, 0, 0
	for r < len(s) {
		in := string(s[r])
		// 需要收缩
		if i, ok := window[in]; ok {
			for ; l <= i; l++ {
				delete(window, string(s[l]))
			}
		}
		window[in] = r
		if res < r-l+1 {
			res = r - l + 1
		}
		r++
	}
	return res
}
