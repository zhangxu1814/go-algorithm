package binaryTree

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	var max, min int
	if p.Val > q.Val {
		max = p.Val
		min = q.Val
	} else {
		max = q.Val
		min = p.Val
	}
	var find func(node *TreeNode) *TreeNode
	find = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val > max {
			return find(node.Left)
		}
		if node.Val < min {
			return find(node.Right)
		}

		return node
	}

	return find(root)
}
