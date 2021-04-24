package queue_stack

import "testing"

func TestQueueStack(t *testing.T) {
	result := []int{}
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	result = append(result, s.Pop())
	if s.Top() != 3 {
		t.Errorf("get: %d, want: %d", s.Top(), 3)
	}
	s.Push(5)
	if s.Top() != 5 {
		t.Errorf("get: %d, want: %d", s.Top(), 5)
	}
	result = append(result, s.Pop())
	result = append(result, s.Pop())
	result = append(result, s.Pop())
	if s.Top() != 1 {
		t.Errorf("get: %d, want: %d", s.Top(), 1)
	}
	result = append(result, s.Pop())

	want := []int{4, 5, 3, 2, 1}
	for i, v := range result {
		if v != want[i] {
			t.Errorf("get: %d, want: %d", v, want[i])
		}
	}
}
