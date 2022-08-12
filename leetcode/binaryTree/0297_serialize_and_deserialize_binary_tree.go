package binaryTree

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	s := []byte(strconv.Itoa(root.Val))
	s = append(append(s, ","...), this.serialize(root.Left)...)
	s = append(append(s, ","...), this.serialize(root.Right)...)

	return string(s)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	var des func() *TreeNode
	des = func() *TreeNode {
		if s[0] == "#" {
			s = s[1:]
			return nil
		}
		i, _ := strconv.Atoi(s[0])
		s = s[1:]
		root := &TreeNode{Val: i}
		root.Left = des()
		root.Right = des()

		return root
	}
	return des()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
