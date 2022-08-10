package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res, depth int

//遍历
func maxDepth(root *TreeNode) int {
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	depth++
	if root.Left == nil && root.Right == nil {
		res = depth
	}
	traverse(root.Left)
	traverse(root.Right)
	depth--
}

//递归
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := maxDepth2(root.Left), maxDepth2(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
