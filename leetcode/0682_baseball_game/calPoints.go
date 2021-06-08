package baseball_game

import (
	"strconv"
)

type Stack struct {
	s []int
}

func (this *Stack) push(x int) {
	this.s = append(this.s, x)
}

func (this *Stack) top() int {
	return this.s[len(this.s)-1]
}

func (this *Stack) pop() int {
	x := this.top()
	this.s = this.s[:len(this.s)-1]
	return x
}

//TODO use switch replace if
func calPoints(ops []string) int {
	stack := Stack{[]int{}}
	result := 0
	for _, v := range ops {
		if v == "+" {
			s1 := stack.pop()
			s2 := stack.top()
			s := s1 + s2
			stack.push(s1)
			stack.push(s)
			result += s
			continue
		}
		if v == "D" {
			s := stack.top()
			stack.push(s * 2)
			result += s * 2
			continue
		}
		if v == "C" {
			s := stack.pop()
			result -= s
			continue
		}
		n, err := strconv.Atoi(v)
		if err == nil {
			stack.push(n)
			result += n
		}
	}

	return result
}
