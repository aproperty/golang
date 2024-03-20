package main

func Revert(src *Node) (new *Node) {

	dummyNode := &Node{0, nil}

	cursor := src

	for cursor != nil {

		p := cursor
		cursor = cursor.next

		p.next = dummyNode.next
		dummyNode.next = p
	}

	return dummyNode.next
}
