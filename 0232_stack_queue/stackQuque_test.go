package stack_queue

import "testing"

func TestStackQueue(t *testing.T) {
	result := []int{}
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	result = append(result, q.Pop())
	q.Push(5)
	result = append(result, q.Pop())
	result = append(result, q.Pop())
	result = append(result, q.Pop())
	result = append(result, q.Pop())

	expect := []int{1, 2, 3, 4, 5}
	for i, v := range result {
		if v != expect[i] {
			t.Errorf("got: %d, want: %d", v, expect[i])
		}
	}
}
