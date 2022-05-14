package main

import "math"

func main() {
	println(eat(2, 2))
}

// 二分查询左边界, k 取值范围 1, 10^9
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1000000000
	for l <= r {
		mid := l + (r-l)>>1
		// l, mid , h,  r
		if eatSpeed(piles, mid) <= h {
			r = mid - 1
		} else if eatSpeed(piles, mid) > h {
			l = mid + 1
		}
	}
	return l
}

// 以速度 k 吃完一堆香蕉的时间
func eatSpeed(piles []int, k int) int {
	time := 0
	for _, pile := range piles {
		time += eat(pile, k)
	}
	return time
}

// eat
func eat(total int, k int) int {
	return int(math.Ceil(float64(total) / float64(k)))
}
