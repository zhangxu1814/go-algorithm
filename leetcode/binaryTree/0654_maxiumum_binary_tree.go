package binaryTree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}

	max, k := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			k = i
		}
	}

	root := &TreeNode{Val: max}
	if k == 0 {
		root.Left = constructMaximumBinaryTree(nil)
	} else {
		root.Left = constructMaximumBinaryTree(nums[:k])
	}
	if k == len(nums)-1 {
		root.Right = constructMaximumBinaryTree(nil)
	} else {
		root.Right = constructMaximumBinaryTree(nums[k+1:])
	}

	return root
}
