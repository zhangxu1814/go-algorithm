package binaryTree

func generateTrees(n int) []*TreeNode {
	var count func(lo, hi int) []*TreeNode
	count = func(lo, hi int) []*TreeNode {
		if lo > hi {
			return []*TreeNode{nil}
		}

		res := make([]*TreeNode, 0)
		for i := lo; i <= hi; i++ {
			left := count(lo, i-1)
			right := count(i+1, hi)
			for l := 0; l < len(left); l++ {
				for r := 0; r < len(right); r++ {
					node := new(TreeNode)
					node.Val = i
					node.Left = left[l]
					node.Right = right[r]
					res = append(res, node)
				}
			}
		}

		return res
	}

	return count(1, n)
}
