package _6

import "math"

// 涵盖 t 的最小子串 -> 滑动窗口， 子串唯一 需要遍历完成判断切片长度
// 纯英文字符 A ~ Z +32 a ~ z 26+32 = 58
func minWindow(s string, t string) string {
	window, need := make(map[string]int, len(t)), make(map[string]int, len(t))
	for _, c := range t {
		need[string(c)]++
	}
	left, right, valid := 0, 0, 0
	start, end := 0, math.MaxInt64
	for right < len(s) {
		c := string(s[right])
		right++
		// 只插入需要的值，减少缩减左边界时的操作
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for valid == len(need) {
			r := string(s[left])
			if _, ok := need[r]; ok {
				if window[r] == need[r] {
					if right-left < end-start {
						start = left
						end = right
					}
					valid--
				}
				window[r]--
			}
			left++
		}

	}
	if end == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}
