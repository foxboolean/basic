package _47

// 遍历以每个元素为中点向两侧遍历
func countSubstrings(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += palindrome(s, i, i)
		sum += palindrome(s, i, i+1)
	}
	return sum
}

func palindrome(s string, l int, r int) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}
