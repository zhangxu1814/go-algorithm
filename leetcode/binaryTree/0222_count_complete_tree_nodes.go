package binaryTree

import "math"

func countNodes(root *TreeNode) int {
	l, r, ln, rn := root, root, 0, 0
	for l != nil {
		ln++
		l = l.Left
	}
	for r != nil {
		rn++
		r = r.Right
	}
	if ln == rn {
		return int(math.Pow(2, float64(ln)) - 1)
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
