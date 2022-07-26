package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{}
	p.Next = head
	p1, p2 := p, p
	for i := 0; i < n+1; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return p.Next
}
