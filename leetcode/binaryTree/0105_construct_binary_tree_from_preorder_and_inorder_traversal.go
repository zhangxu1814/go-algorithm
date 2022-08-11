package binaryTree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}

	h, k := preorder[0], -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == h {
			k = i
			break
		}
	}

	root := &TreeNode{Val: h}
	if k == 0 {
		root.Left = buildTree(nil, nil)
	} else {
		root.Left = buildTree(preorder[1:k+1], inorder[:k])
	}
	if k == len(inorder)-1 {
		root.Right = buildTree(nil, nil)
	} else {
		root.Right = buildTree(preorder[k+1:], inorder[k+1:])
	}

	return root
}
