package min_stack

type MinStack struct {
	stack []item
}

type item struct {
	val, min int
}

func Constructor() MinStack {
	return MinStack{
		stack: []item{},
	}
}

func (this *MinStack) push(val int) {
	min := val
	if len(this.stack) > 0 && this.GetMin() < min {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{val: val, min: min})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
