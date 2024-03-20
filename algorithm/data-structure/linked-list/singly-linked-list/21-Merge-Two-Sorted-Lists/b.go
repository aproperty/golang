package main

// 取两个有序链表中较小的那一个数，链接到新链表上，
// 直到其中一个链表为空。再把另一个链表的剩余部分，链接到新链表上。
func mergeTwoListsBb(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var l3 = &ListNode{}

	var current = l3
	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return l3.Next
}
