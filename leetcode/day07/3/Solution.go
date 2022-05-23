package main

// 无重复最长子串，<字符，索引>
func lengthOfLongestSubstring(s string) int {
	need := make(map[string]int, len(s))
	res, l, r := 0, 0, 0
	for r < len(s) {
		in := string(s[r])
		if i, ok := need[in]; ok {
			for ; l <= i; l++ {
				delete(need, string(s[l]))
			}
		}
		need[in] = r
		if res < r-l+1 {
			res = r - l + 1
		}
		r++
	}
	return res
}
