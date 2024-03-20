package main

func inorderTraversalBb(root *TreeNode) (res []int) {

	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		res = append(res, root.Val)

		root = root.Right
	}

	return
}
