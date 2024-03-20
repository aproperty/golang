package main

type LRUCache struct {
	capacity int
	len      int

	hashMap map[int]*Node

	head *Node
	tail *Node
}

type Node struct {
	Key int
	Val int

	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {

	m := make(map[int]*Node)

	lru := LRUCache{
		capacity: capacity,
		hashMap:  m,
		head:     &Node{},
		tail:     &Node{},
	}

	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head

	return lru
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);

 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
