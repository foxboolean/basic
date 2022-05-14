package main

// 双指针处理回文
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// [) 左闭右开
	return s[l+1 : r]
}

func main() {
	i := palindrome("babad", 2, 2)
	println(i)
}
