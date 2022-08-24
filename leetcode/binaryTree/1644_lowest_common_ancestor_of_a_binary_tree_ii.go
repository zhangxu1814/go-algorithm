package binaryTree

func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	findP, findQ := false, false
	var find func(node *TreeNode, p, q int) *TreeNode
	find = func(node *TreeNode, p, q int) *TreeNode {
		if node == nil {
			return nil
		}

		l := find(node.Left, p, q)
		r := find(node.Right, p, q)

		if l != nil && r != nil {
			return node
		}

		if node.Val == p {
			findP = true
			return node
		}
		if node.Val == q {
			findQ = true
			return node
		}

		if l != nil {
			return l
		} else {
			return r
		}
	}

	res := find(root, p.Val, q.Val)
	if findP && findQ {
		return res
	}
	return nil
}
