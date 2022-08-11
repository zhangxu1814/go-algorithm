package binaryTree

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || postorder == nil {
		return nil
	}

	h, k := postorder[len(postorder)-1], -1
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
		root.Left = buildTree(inorder[:k], postorder[:k])
	}
	if k == len(inorder)-1 {
		root.Right = buildTree(nil, nil)
	} else {
		root.Right = buildTree(inorder[k+1:], postorder[k:len(postorder)-1])
	}

	return root
}
