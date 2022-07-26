package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	p := &ListNode{}
//	p0, p1, p2 := p, list1, list2
//	for {
//		if p1 == nil || p2 == nil {
//			break
//		}
//
//		if p1.Val > p2.Val {
//			p0.Next = p2
//			p2 = p2.Next
//		} else {
//			p0.Next = p1
//			p1 = p1.Next
//		}
//		p0 = p0.Next
//	}
//
//	if p1 == nil {
//		p0.Next = p2
//	}
//	if p2 == nil {
//		p0.Next = p1
//	}
//
//	return p.Next
//}

// 递归
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
