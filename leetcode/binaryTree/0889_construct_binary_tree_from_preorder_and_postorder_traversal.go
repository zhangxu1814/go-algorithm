package binaryTree

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	h, h1, k := preorder[0], preorder[1], -1
	for i := 0; i < len(postorder); i++ {
		if postorder[i] == h1 {
			k = i
			break
		}
	}

	root := &TreeNode{Val: h}
	root.Left = constructFromPrePost(preorder[1:k+2], postorder[:k+1])
	root.Right = constructFromPrePost(preorder[k+2:], postorder[k+1:len(postorder)-1])

	return root
}
