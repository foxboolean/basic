package main

func main() {
	inclusion := checkInclusion("hello",
		"ooolleoooleh")
	println(inclusion)
}

// s2 是否包含 s1「排列」 的子序列
func checkInclusion(s1 string, s2 string) bool {
	need, window := [26]int{}, [26]int{} // 两个长度为26的数组

	for i := 0; i < len(s1); i++ { // 统计s1的字符的出现次数
		need[s1[i]-'a']++
	}
	start := 0 // 指向s2子串的开头

	for i, c := range s2 { // 遍历s2的字符
		window[c-'a']++ // 将s2的字符统计到window

		// 开启循环，在start没有越界的前提下，start上的字符，window的比need多
		// 这说明start不是合格子串的开头，start应该步进，舍弃这个字符
		for start <= i && need[s2[start]-'a'] < window[s2[start]-'a'] {
			window[s2[start]-'a']--
			start++
		}
		// 如果两个count数组全等，说明找到了满足条件的子串
		if need == window { // go里面数组是基本类型，可以这么比较
			return true
		}
	}

	// 考察了所有s2的字符，都没有找到合适的子串
	return false
}
