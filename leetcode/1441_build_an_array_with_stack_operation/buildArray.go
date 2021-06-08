package build_an_array_with_stack_operation

func buildArray(target []int, n int) []string {
	l := len(target)
	p := 0
	result := []string{}
	for i := 1; i <= n; i++ {
		if p >= l {
			break
		} else {
			if target[p] == i {
				result = append(result, "Push")
				p++
			} else {
				result = append(result, "Push")
				result = append(result, "Pop")
			}
		}
	}
	return result
}

//TODO p++ can be replaced by slice operation
