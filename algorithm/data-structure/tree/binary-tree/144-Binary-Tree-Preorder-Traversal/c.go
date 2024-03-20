package main

// Morris 其实 解决了 一个常规循环中 循环到 叶子节点后 难以回到 根节点的问题。
// 我们 都知道 前序遍历 是 先左后右，那么对任一节点 p1 来说，其右子树 p1.right 所有节点 必然在左子树 p1.left之后。

// 代码中 第二个 while 做的是，在 p1.left 里一直往右，直到找不到更右的点，记这一点为 p2。
// 然后把 p1.right 接到 p2 的右边。 这样既保证了 p1.right 在 p1.left 所有点之后，又不需要再回到 p1 节点。
// 即在正常的 往下循环 的过程中，不断把右半部分 剪下来，接到 左半部分的 最右下

// Morris 遍历
func preorderTraversalCc(root *TreeNode) (vals []int) {

	var p1 *TreeNode = root
	var p2 *TreeNode = nil

	for p1 != nil {

		p2 = p1.Left
		if p2 != nil { // 如果 当前节点 的 左子节点 不为空，

			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right // 在 当前节点 的 左子树中      找到 当前节点 在中序遍历下 的前驱节点：
			}

			if p2.Right == nil { // 如果 前驱节点 的 右子节点 为空，

				vals = append(vals, p1.Val) // 将 当前节点 加入答案，

				p2.Right = p1 // 将 前驱节点 的 右子节点 设置为 当前节点。

				p1 = p1.Left // 当前节点 更新为 当前节点 的 左子节点。

				continue
			}

			// 如果 前驱节点 的 右子节点 为 当前节点，
			p2.Right = nil // 将它的 右子节点 重新设为空。

		} else { // 如果当前节点的左子节点为空，将当前节点加入答案，并遍历当前节点的右子节点；
			vals = append(vals, p1.Val)
		}

		p1 = p1.Right // 当前节点 更新为 当前节点 的 右子节点。
	}

	return
}
