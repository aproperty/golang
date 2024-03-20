package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// 递归出口

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 递归逻辑
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
