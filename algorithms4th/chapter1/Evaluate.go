package algorithms4th

import (
	"fmt"
	"math"
	"strconv"

	"github.com/zhangxu1814/algorithm/structures"
)

func evaluate(s []string) {
	ops := structures.NewStack()
	vals := structures.NewStack()

	for _, v := range s {
		switch v {
		case "(":
		case "+":
			ops.Push(v)
		case "-":
			ops.Push(v)
		case "*":
			ops.Push(v)
		case "/":
			ops.Push(v)
		case "sqrt":
			ops.Push(v)
		case ")":
			op := ops.Pop()
			val := vals.Pop().(float64)
			switch op {
			case "+":
				val = vals.Pop().(float64) + val
			case "-":
				val = vals.Pop().(float64) - val
			case "*":
				val = vals.Pop().(float64) * val
			case "/":
				val = vals.Pop().(float64) / val
			case "sqrt":
				val = math.Sqrt(val)
			}
			vals.Push(val)
		default:
			val, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println("err : ", err)
			}
			vals.Push(val)
		}
	}

	fmt.Println("result is : ", vals.Pop())
}
