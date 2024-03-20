package heapsort

func maxHeap(root int, end int, c []int) {

	for {
		var child = 2*root + 1

		// 判断是否存在 child 节点
		if child > end {
			break
		}

		// 判断 右 child 是否存在，如果存在 则 和另外一个同级节点 进行比较
		if child+1 <= end && c[child] < c[child+1] {
			child += 1
		}

		if c[root] < c[child] {

			c[root], c[child] = c[child], c[root]
			root = child

		} else {
			break
		}
	}
}
