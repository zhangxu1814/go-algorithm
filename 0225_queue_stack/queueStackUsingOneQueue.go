package queue_stack

type Stack1 struct {
	queue queue
}

func Constructor1() Stack1 {
	return Stack1{queue{}}
}

func (this *Stack1) Push1(x int) {
	this.queue.push(x)
}

func (this *Stack1) Pop1() int {
	l := len(this.queue.val)
	for i := 0; i < l-1; i++ {
		v := this.queue.pop()
		this.queue.push(v)
	}
	v := this.queue.pop()

	return v
}

func (this *Stack1) Top1() int {
	l := len(this.queue.val)
	for i := 0; i < l-1; i++ {
		v := this.queue.pop()
		this.queue.push(v)
	}
	v := this.queue.pop()
	this.queue.push(v)

	return v
}

func (this *Stack1) Empty1() bool {
	return this.queue.empty()
}
