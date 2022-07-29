package linkedlist

import "testing"

func Test234(t *testing.T) {
	head := &ListNode{1, nil}
	head1 := &ListNode{1, nil}
	head2 := &ListNode{2, nil}
	head3 := &ListNode{1, nil}
	head.Next = head1
	head1.Next = head2
	head2.Next = head3
	isPalindrome(head)
}
