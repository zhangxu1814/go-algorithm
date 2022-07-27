package linkedlist

//递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//迭代
func reverse(head *ListNode) *ListNode {
	p, tmp := head, &ListNode{}
	var pre *ListNode

	for p != nil {
		tmp = p.Next
		p.Next = pre
		pre, p = p, tmp
	}

	return pre
}
