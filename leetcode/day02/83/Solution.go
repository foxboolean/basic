package main

// 有序，去重
// 快慢指针，slow 指向「重复的节点」
func deleteDuplicates(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow = slow.Next
			slow.Val = fast.Val
		}
		fast = fast.Next
	}
	if slow != nil {
		slow.Next = nil
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
