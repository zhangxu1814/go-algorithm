package remove_all_adjacent_duplicates_in_string

func removeDuplicates(S string) string {
	stack := []byte{}
	for _, v := range []byte(S) {
		if len(stack) != 0 && v == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, v)
	}
	return string(stack)
}
