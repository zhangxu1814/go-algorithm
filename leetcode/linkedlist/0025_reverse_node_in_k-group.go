package linkedlist

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	newHead := reverseAB(head, cur)
	nxt := reverseKGroup(cur, k)
	head.Next = nxt
	return newHead
}

func reverseN1(head *ListNode, n int) *ListNode {
	var pre *ListNode
	cur, nxt := head, head
	for i := 0; i < n; i++ {
		nxt = cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}
	return pre
}

func reverseAB(a, b *ListNode) *ListNode {
	var pre, cur, nxt *ListNode = nil, a, a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}
	return pre
}
