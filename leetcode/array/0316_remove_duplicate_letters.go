package array

func removeDuplicateLetters(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	res, in := make([]byte, 0), make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		m[s[i]]--

		if in[s[i]] {
			continue
		}

		for len(res) > 0 && res[len(res)-1] > s[i] {
			e := res[len(res)-1]
			if m[e] == 0 {
				break
			}
			res = res[:len(res)-1]
			in[e] = false
		}
		res = append(res, s[i])
		in[s[i]] = true
	}

	return string(res)
}
