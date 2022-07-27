package linkedlist

//递归
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var tmp *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		tmp = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = tmp
	return last
}

//迭代
func reverseD(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	var pre, l, r, suc *ListNode
	p := dummy
	for i := 0; i < left-1; i++ {
		p = p.Next
	}
	pre = p
	p = p.Next
	l = p
	r = p
	p = p.Next
	for i := left; i < right; i++ {
		suc = p.Next
		p.Next = r
		r, p = p, suc
	}
	pre.Next = r
	l.Next = suc

	return dummy.Next
}
