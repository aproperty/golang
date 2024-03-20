package main

import "fmt"

// 递归
func preorderTraversalAa(root *TreeNode) []int {

	var result []int

	var preOrder func(*TreeNode)

	preOrder = func(root *TreeNode) {

		if root == nil {
			return
		}

		result = append(result, root.Val)
		fmt.Println(result)

		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)

	return result
}

// func preorderTraversal(root *TreeNode) (vals []int) {

//     var preorder func(*TreeNode)

//     preorder = func(node *TreeNode) {
//         if node == nil {
//             return
//         }
//         vals = append(vals, node.Val)
//         preorder(node.Left)
//         preorder(node.Right)
//     }

//     preorder(root)

//     return
// }
