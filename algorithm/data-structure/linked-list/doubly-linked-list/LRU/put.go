package main

func (this *LRUCache) Put(key int, value int) {

	if v, ok := this.hashMap[key]; ok {

		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev

		n := this.head.Next

		this.head.Next = v
		v.Prev = this.head

		n.Prev = v
		v.Next = n

		v.Val = value

		return
	}

	if this.len < this.capacity {

		this.len++

		node := &Node{
			Key: key,
			Val: value,
		}
		this.hashMap[key] = node

		n := this.head.Next

		this.head.Next = node
		node.Prev = this.head

		node.Next = n
		n.Prev = node

	} else {

		t := this.tail.Prev

		this.tail.Prev.Prev.Next = this.tail
		this.tail.Prev = this.tail.Prev.Prev

		t.Val = value
		delete(this.hashMap, t.Key)
		t.Key = key

		this.hashMap[key] = t

		hn := this.head.Next

		this.head.Next = t
		t.Prev = this.head

		t.Next = hn
		hn.Prev = t
	}

	return
}
