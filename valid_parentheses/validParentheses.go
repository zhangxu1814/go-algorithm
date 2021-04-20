package valid_parentheses

func isValid(s string) bool {
	l := len(s)
	stack := make([]byte, 0, l)

	for i := 0; i < l; i++ {
		switch v := s[i]; v {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			//pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != v {
				return false
			}
		}
	}

	return len(stack) == 0
}
