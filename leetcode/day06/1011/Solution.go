package main

func main() {
	time := shipWithinDays([]int{1, 2, 3, 1, 1}, 4)
	//time := transitTime(2, []int{1, 2, 3, 1, 1})
	println(time)
}

// 运载和时间是单调递减函数，二分查询
func shipWithinDays(weights []int, days int) int {
	l, r := minCapacity(weights), 25000000
	for l <= r {
		mid := l + (r-l)>>1
		// l,  day , mid , r
		if transitTime(mid, weights) <= days {
			r = mid - 1
		} else if transitTime(mid, weights) > days {
			l = mid + 1
		}
	}
	return l
}

// 运载能力为 capacity 的船运输 weights 获取需要的时间
func transitTime(capacity int, weights []int) int {
	time := 0
	sum := 0
	for i := range weights {
		sum += weights[i]
		if i < len(weights)-1 && sum+weights[i+1] > capacity {
			sum = 0
			time++
		}
	}
	return time + 1
}

func minCapacity(weights []int) int {
	max := 0
	for _, weight := range weights {
		if max < weight {
			max = weight
		}
	}
	return max
}
