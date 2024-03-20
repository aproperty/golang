package main

// 给定一个单向链表，求出该链表中倒数第 n 个元素。
func getItem(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	if n == 0 { // 如果 n=0，则链表的最后一个元素就是所求的元素
		n = 1
	}

	first := head
	for i := 0; i < n-1; i++ { // 第一个指针 从 链表 头指针 开始遍历，向前走 n-1 步，第二个指针不动。
		if first.Next != nil {
			first = first.Next
		} else {
			return nil
		}
	}

	// 从第 n 步 开始，第二个指针 也开始 从 链表 头指针 开始遍历
	// 由于 两个指针 之间间隔 保持 n-1，
	// 当第一个指针 到达链表的 尾节点时，第二个指针 指向的 正好是 倒数 第 n 个节点
	second := head
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	return second
}
