package remove_outermost_parentheses

func removeOutermostParenthese(S string) string {
	stack := []byte{}
	count := []int{}
	for i, v := range []byte(S) {
		if v == '(' {
			stack = append(stack, v)
			continue
		}
		if top := stack[len(stack)-1]; top == '(' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				count = append(count, i)
			}
		}
	}
	//TODO can optimize the string operation
	start := 0
	result := []byte{}
	for _, v := range count {
		if v-start > 1 {
			result = append(result, []byte(S)[start+1:v]...)
		}
		start = v + 1
	}

	return string(result)
}
