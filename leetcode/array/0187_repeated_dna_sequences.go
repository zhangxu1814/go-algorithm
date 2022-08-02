package array

func findRepeatedDnaSequences(s string) []string {
	l, r := 0, 10
	rep := make(map[string]bool)
	res := make([]string, 0)
	for r <= len(s) {
		if added, exi := rep[s[l:r]]; exi {
			if !added {
				res = append(res, s[l:r])
				rep[s[l:r]] = true
			}
		} else {
			rep[s[l:r]] = false
		}
		l++
		r++
	}

	return res
}
