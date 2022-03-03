package main

import "fmt"

func checkInclusion(s2 string, s1 string) bool {
	need := make(map[string]int, len(s2))
	for i := 0; i < len(s2); i++ {
		need[string(s2[i])]++
	}

	left, right, valid := 0, 0, 0
	window := make(map[string]int, len(s1))
	for right < len(s1) {
		in := string(s1[right])
		right++
		if need[in] > 0 {
			if left == -1 {
				left = right - 1
			}
			if window[in] < need[in] {
				window[in]++
				if window[in] == need[in] {
					valid++
					if valid == len(need) {
						return true
					}
				}
			} else {
				for i := left; i < right; i++ {
					out := string(s1[left])
					if out == in {
						left++
						break
					}
					window[out]--
					if window[out] == 0 {
						valid--
					}
				}
			}

		} else {
			left = -1
			valid = 0
		}
	}
	return false
}

func main() {
	res := checkInclusion("adc", "dcda")
	fmt.Println(res)
}
