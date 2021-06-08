package queue_stack

type Stack struct {
	in  *queue
	out *queue
}

type queue struct {
	val []int
}

func (this *queue) push(x int) {
	this.val = append(this.val, x)
}

func (this *queue) pop() int {
	v := this.top()
	this.val = this.val[1:]
	return v
}

func (this *queue) top() int {
	return this.val[0]
}

func (this *queue) empty() bool {
	return len(this.val) == 0
}

func Constructor() Stack {
	in := new(queue)
	out := new(queue)
	return Stack{in, out}
}

func (this *Stack) Push(x int) {
	this.in.push(x)
}

func swap(s *Stack) {
	tmp := s.in
	s.in = s.out
	s.out = tmp
}

func (this *Stack) Pop() int {
	l := len(this.in.val)
	for i := 0; i < l-1; i++ {
		v := this.in.pop()
		this.out.push(v)
	}
	v := this.in.pop()
	swap(this)

	return v
}

func (this *Stack) Top() int {
	l := len(this.in.val)
	for i := 0; i < l-1; i++ {
		v := this.in.pop()
		this.out.push(v)
	}
	v := this.in.pop()
	this.out.push(v)
	swap(this)

	return v
}

func (this *Stack) Empty() bool {
	return this.in.empty()
}
