package queue

type Queue struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	front *node
	back  *node
	value interface{}
}

func Constructor() *Queue {
	return &Queue{nil, nil, 0}
}

func (this *Queue) push(v interface{}) {
	n := &node{nil, nil, v}
	if this.length == 0 {
		this.head = n
		this.tail = this.head
	} else {
		n.front = this.tail
		this.tail.back = n
		this.tail = n
	}
	this.length++
}

func (this *Queue) pop() interface{} {
	if this.length == 0 {
		return nil
	}

	v := this.head.value
	if this.head.back == nil {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.back
		this.head.front.back = nil
		this.head.front = nil
	}
	this.length--

	return v
}

func (this *Queue) peek() interface{} {
	if this.head == nil {
		return nil
	}
	return this.head.value
}

func (this *Queue) isEmpty() bool {
	return !(this.length > 0)
}
