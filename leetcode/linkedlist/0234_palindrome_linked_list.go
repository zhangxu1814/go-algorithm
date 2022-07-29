package linkedlist

var h *ListNode

func isPalindrome(head *ListNode) bool {
	h = head
	return traverse(head)
}

func traverse(head *ListNode) bool {
	if head == nil {
		return true
	}
	res := traverse(head.Next)
	res = res && head.Val == h.Val
	h = h.Next
	return res
}

func isPalindrome2(head *ListNode) bool {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//反转
	var pre *ListNode
	cur := slow.Next
	nxt := cur
	for cur != nil {
		nxt, cur.Next = nxt.Next, pre
		pre, cur = cur, nxt
	}
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre, head = pre.Next, head.Next
	}
	return true
}
