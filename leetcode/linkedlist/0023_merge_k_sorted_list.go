package linkedlist

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	tmp := &ListNode{}
	p := tmp

	for _, l := range lists {
		if l != nil {
			heap.Push(pq, l)
		}

	}
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*ListNode)
		if item.Next != nil {
			heap.Push(pq, item.Next)
		}
		p.Next = item
		p = p.Next
	}

	return tmp.Next
}

//检查是否实现了接口
var _ heap.Interface = &PriorityQueue{}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	item := old[len(old)-1]
	old[len(old)-1] = nil
	*pq = old[:len(old)-1]
	return item
}

//TODO 分治策略
