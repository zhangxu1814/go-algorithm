package binaryTree

func kthSmallest(root *TreeNode, k int) int {
	var rank, res int
	var traverseBST func(*TreeNode)
	traverseBST = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverseBST(node.Left)
		rank++
		if rank == k {
			res = node.Val
			return
		}
		if rank > k {
			return
		}
		traverseBST(node.Right)
	}
	traverseBST(root)

	return res
}
