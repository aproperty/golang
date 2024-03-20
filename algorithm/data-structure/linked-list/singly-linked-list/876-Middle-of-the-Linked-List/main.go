package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路是快慢指针，快指针比慢指针的步长快一步，那么快指针跑到链表末尾的时候，慢指针此时就位于链表的中间位置。
func middleNode(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
