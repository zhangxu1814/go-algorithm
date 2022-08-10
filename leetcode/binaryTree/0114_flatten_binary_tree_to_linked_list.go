package binaryTree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	cur, tmp := root, root.Right
	root.Left, root.Right = nil, root.Left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = tmp
}
