package binaryTree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var find func(node *TreeNode, a, b int) *TreeNode
	find = func(node *TreeNode, a, b int) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == a || node.Val == b {
			return node
		}
		l := find(node.Left, a, b)
		r := find(node.Right, a, b)
		if l != nil && r != nil {
			return node
		}

		if l != nil {
			return l
		} else {
			return r
		}
	}

	return find(root, p.Val, q.Val)
}
