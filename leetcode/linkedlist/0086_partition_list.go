package linkedlist

func partition(head *ListNode, x int) *ListNode {
	small, big, p := &ListNode{}, &ListNode{}, head
	ps, pb := small, big
	for {
		if p == nil {
			break
		}
		if p.Val < x {
			ps.Next = p
			ps = ps.Next
		} else {
			pb.Next = p
			pb = pb.Next
		}
		p = p.Next
	}
	pb.Next = nil
	ps.Next = big.Next
	return small.Next
}
