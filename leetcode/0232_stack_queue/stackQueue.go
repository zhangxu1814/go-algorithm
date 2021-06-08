package stack_queue

type Queue struct {
	in  []int
	out []int
}

func Constructor() Queue {
	return Queue{nil, nil}
}

func (this *Queue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *Queue) Pop() int {
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	v := this.out[len(this.out)-1]

	this.in = this.in[1:]
	this.out = nil

	return v
}

func (this *Queue) Peek() int {
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	v := this.out[len(this.out)-1]

	this.out = nil

	return v
}

func (this *Queue) Empty() bool {
	return len(this.in) == 0
}
