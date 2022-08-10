package binaryTree

//遍历
//func invertTree(root *TreeNode) *TreeNode {
//	traverse1(root)
//	return root
//}
//
//func traverse1(root *TreeNode) {
//	if root == nil {
//		return
//	}
//
//	tmp := root.Left
//	root.Left = root.Right
//	root.Right = tmp
//
//	traverse1(root.Left)
//	traverse1(root.Right)
//}

//递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
