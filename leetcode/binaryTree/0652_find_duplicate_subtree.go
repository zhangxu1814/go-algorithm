package binaryTree

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	sub := make(map[string]int)
	list := make([]*TreeNode, 0)

	var find func(node *TreeNode) string
	find = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}

		s := strconv.Itoa(node.Val) + "," + find(node.Left) + "," + find(node.Right)
		sub[s]++
		if sub[s] == 2 {
			list = append(list, node)
		}
		return s
	}
	find(root)

	return list
}
