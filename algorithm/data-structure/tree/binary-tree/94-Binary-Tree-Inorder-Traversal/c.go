package main

// Morris 中序遍历

// 假设当前遍历到的节点为 x

func inorderTraversalCc(root *TreeNode) (res []int) {

	for root != nil {

		if root.Left != nil {

			predecessor := root.Left

			for predecessor.Right != nil && predecessor.Right != root { // 有右子树 且 没有设置过 指向 root，则继续向右走
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left

			} else {
				res = append(res, root.Val)
				predecessor.Right = nil
				root = root.Right
			}

		} else { // 如果没有左子树
			res = append(res, root.Val)
			root = root.Right
		}

	} //

	return
}
