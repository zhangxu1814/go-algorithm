package binaryTree

func isValidBST(root *TreeNode) bool {
	cur := -1<<31 - 1
	res := true
	var trav func(node *TreeNode)
	trav = func(node *TreeNode) {
		if node == nil || res == false {
			return
		}
		trav(node.Left)
		if node.Val <= cur {
			res = false
			return
		}
		cur = node.Val
		trav(node.Right)
	}

	trav(root)

	return res
}
