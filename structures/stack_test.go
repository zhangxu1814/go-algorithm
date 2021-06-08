package structures

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	result := []int{}
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	result = append(result, s.Pop().(int))
	if s.Top() != 3 {
		t.Errorf("get: %d, want: %d", s.Top(), 3)
	}
	s.Push(5)
	if s.Top() != 5 {
		t.Errorf("get: %d, want: %d", s.Top(), 5)
	}
	result = append(result, s.Pop().(int))
	result = append(result, s.Pop().(int))
	result = append(result, s.Pop().(int))
	if s.Top() != 1 {
		t.Errorf("get: %d, want: %d", s.Top(), 1)
	}
	result = append(result, s.Pop().(int))
}
