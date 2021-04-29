package structures

type ListNode struct {
	value interface{}
	next  *ListNode
}

type DoublyListNode struct {
	prev  *DoublyListNode
	next  *DoublyListNode
	value interface{}
}
