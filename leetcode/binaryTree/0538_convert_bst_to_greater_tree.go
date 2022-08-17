package binaryTree

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var conv func(node *TreeNode)
	conv = func(node *TreeNode) {
		if node == nil {
			return
		}
		conv(node.Right)
		sum += node.Val
		node.Val = sum
		conv(node.Left)
	}

	conv(root)
	return root
}
