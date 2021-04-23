package longest_substring

func longestSubstring(s string) int {
	m := make(map[rune]int)
	count, head, tail := 0, 0, 0
	for i, val := range s {
		if v, exist := m[val]; exist && v >= head {
			tail++
			head = v + 1
			m[val] = i
			count = max(count, tail-head)
			continue
		}
		m[val] = i
		tail++
		count = max(count, tail-head)
	}

	return count
}

func max(m int, n int) int {
	if m >= n {
		return m
	} else {
		return n
	}
}

//TODO
// func longestSubstringUsingQueue(s string) {

// }
