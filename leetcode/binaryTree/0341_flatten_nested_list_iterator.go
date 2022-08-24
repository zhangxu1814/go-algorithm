package binaryTree

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	it       []int
	len, cur int
}

//func Constructor(nestedList []*NestedInteger) *NestedIterator {
//	it := make([]int, 0)
//	var traverse func(node *NestedInteger)
//	traverse = func(node *NestedInteger) {
//		if node.IsInteger() {
//			it = append(it, node.GetInteger())
//		}
//		for _, n := range node.GetList() {
//			traverse(n)
//		}
//	}
//	for _, n := range nestedList {
//		traverse(n)
//	}
//	return &NestedIterator{
//		it:  it,
//		len: len(it),
//		cur: -1,
//	}
//}

func (this *NestedIterator) Next() int {
	this.cur++
	return this.it[this.cur]
}

func (this *NestedIterator) HasNext() bool {
	if this.cur < this.len-1 {
		return true
	}
	return false
}
