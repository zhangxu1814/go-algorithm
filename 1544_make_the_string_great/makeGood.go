package make_the_string_great

import "math"

func makeGood(s string) string {
	stack := []rune{}
	for _, v := range s {
		if len(stack) > 0 && math.Abs(float64(v-stack[len(stack)-1])) == 32 {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, v)
	}
	return string(stack)
}

//TODO byte:uint8  rune:int32  using unsigned int for subtraction will wrong
