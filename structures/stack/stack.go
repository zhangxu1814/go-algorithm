package structures

type stack struct {
	values []interface{}
	height int
}

func NewStack() *stack {
	return &stack{
		values: []interface{}{},
		height: 0,
	}
}

func (this *stack) Push(value interface{}) {
	this.values = append(this.values, value)
	this.height++
}

func (this *stack) Pop() interface{} {
	if this.Empty() {
		return nil
	}
	top := this.Top()
	this.values = this.values[:this.height-1]
	this.height--
	return top
}

func (this *stack) Top() interface{} {
	if this.Empty() {
		return nil
	}
	return this.values[this.height-1]
}

func (this *stack) Empty() bool {
	return this.height == 0
}
