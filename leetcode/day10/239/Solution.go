package _39

// 最大值单调队列解题
func maxSlidingWindow(nums []int, k int) []int {
	var s []int
	m := MonotonicQueue{S: s}
	var res []int
	for i, num := range nums {
		m.push(num)
		if i >= k-1 {
			// 记录结果
			res = append(res, m.max())
			// 移除左侧元素
			m.pop(nums[i-k+1])
		}
	}
	return res
}

// MonotonicQueue 单调队列，元素从大到小
type MonotonicQueue struct {
	S []int
}

// push 向单调队列中添加元素
func (m *MonotonicQueue) push(i int) {
	// 弹出小元素
	for len(m.S) > 0 && m.S[len(m.S)-1] < i {
		m.S = m.S[:len(m.S)-1]
	}
	m.S = append(m.S, i)
}

// pop 弹出头元素
func (m *MonotonicQueue) pop(n int) {
	if n == m.max() {
		m.S = m.S[1:len(m.S)]
	}
}

// max 返回最大元素
func (m *MonotonicQueue) max() int {
	return m.S[0]
}
