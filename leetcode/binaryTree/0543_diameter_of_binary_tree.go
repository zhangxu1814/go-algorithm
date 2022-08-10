package binaryTree

var m int

func diameterOfBinaryTree(root *TreeNode) int {
	m = 0
	max(root)
	return m
}

func max(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := max(root.Left)
	r := max(root.Right)
	if tmp := l + r; tmp > m {
		m = tmp
	}

	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
