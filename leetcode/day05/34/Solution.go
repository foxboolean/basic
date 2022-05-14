package _4

// 升序，查询范围，二分 -> 求左边界，求右边界
// l mid target r    l target mid r
func searchRange(nums []int, target int) []int {
	l := searchLeft(nums, target)
	r := searchRight(nums, target)
	return []int{l, r}
}

func searchLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] >= target {
			r = mid - 1
		}
	}
	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

func searchRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] <= target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	if r < 0 || nums[r] != target {
		return -1
	}
	return r
}
