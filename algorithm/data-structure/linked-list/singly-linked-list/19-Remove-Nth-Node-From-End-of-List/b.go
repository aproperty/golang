package main

func removeNthFromEndBb(head *ListNode, n int) *ListNode {

	dummy := &ListNode{0, head}

	nodes := []*ListNode{}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next

	return dummy.Next
}
