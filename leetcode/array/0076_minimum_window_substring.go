package array

func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	l, r, matched, min, start := 0, 0, 0, 100000, 0
	for r < len(s) {
		e := s[r]
		r++
		if _, exi := need[e]; exi {
			window[e]++
			if window[e] == need[e] {
				matched++
			}
		}

		for matched == len(need) {
			if r-l < min {
				min = r - l
				start = l
			}
			d := s[l]
			l++
			if _, exi := need[d]; exi {
				if window[d] == need[d] {
					matched--
				}
				window[d]--
			}
		}
	}

	if min == 100000 {
		return ""
	}
	return s[start : start+min]
}
