package main

func RevertBb(src *Node) (new *Node) {

	nodeMap := make(map[int]*Node)

	nextNode := src
	count := 0
	nodeMap[count] = nextNode

	for {
		if nextNode.next != nil {
			nextNode = nextNode.next
			count += 1
			nodeMap[count] = nextNode
		} else {
			break
		}
	}

	for n := count; n > 0; n-- {
		nodeMap[n].next = nodeMap[n-1]
	}
	nodeMap[0].next = nil

	return nodeMap[count]
}
