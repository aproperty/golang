package main

// 迭代
func preorderTraversal(root *TreeNode) (vals []int) {

	stack := []*TreeNode{}

	node := root

	for node != nil || len(stack) > 0 {

		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)

			node = node.Left
		}

		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return
}
