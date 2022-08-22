package binaryTree

func lowestCommonAncestorIV(root *TreeNode, nodes []*TreeNode) *TreeNode {
	values := make(map[int]int)
	for i := 0; i < len(nodes); i++ {
		values[nodes[i].Val]++
	}
	var find func(node *TreeNode) *TreeNode
	find = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if values[node.Val] == 1 {
			return node
		}
		l := find(node.Left)
		r := find(node.Right)
		if l != nil && r != nil {
			return node
		}
		if l != nil {
			return l
		} else {
			return r
		}
	}

	return find(root)
}
