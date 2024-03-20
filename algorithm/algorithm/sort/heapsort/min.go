package heapsort

func minHeap(root int, end int, c []int) {
	for {
		var child = 2*root + 1
		// 判断是否存在child节点
		if child > end {
			break
		}
		// 判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}
