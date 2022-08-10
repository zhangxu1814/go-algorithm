package binaryTree

var res1 []int

func preorderTraversal(root *TreeNode) []int {
	res1 = nil
	traversal(root)
	return res1
}

func traversal(root *TreeNode) {
	if root == nil {
		return
	}

	res1 = append(res1, root.Val)
	traversal(root.Left)
	traversal(root.Right)
}
