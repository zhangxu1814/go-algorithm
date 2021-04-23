package queue

import "testing"

func TestQueue(t *testing.T) {
	q := Constructor()
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	q.pop()
	q.pop()
	q.pop()
	q.pop()
}
