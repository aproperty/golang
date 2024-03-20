package main

import "log"

type Node struct {
	data int
	next *Node
}

func PrintLinkNode(node *Node) {
	if node == nil {
		return
	}
	nextNode := node
	log.Println("PrintLinkNode:")
	for {
		println(nextNode.data)
		if nextNode.next == nil {
			break
		}
		nextNode = nextNode.next
	}
}

func main() {
	first := &Node{data: 1}
	second := &Node{data: 2}
	third := &Node{data: 3}
	four := &Node{data: 4, next: nil}

	first.next = second
	second.next = third
	third.next = four

	PrintLinkNode(first)

	netHead := Revert(first)

	PrintLinkNode(netHead)
}
