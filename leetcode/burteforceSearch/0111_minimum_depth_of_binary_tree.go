package burteforceSearch

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 1

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			q := queue[0]
			queue = queue[1:]
			if q.Left == nil && q.Right == nil {
				return depth
			}
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		depth++
	}

	return depth
}
