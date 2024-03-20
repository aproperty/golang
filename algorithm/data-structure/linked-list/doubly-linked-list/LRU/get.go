package main

func (this *LRUCache) Get(key int) int {

	if v, ok := this.hashMap[key]; ok {

		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev

		n := this.head.Next

		this.head.Next = v
		v.Prev = this.head

		n.Prev = v
		v.Next = n

		return v.Val
	}

	return -1
}
