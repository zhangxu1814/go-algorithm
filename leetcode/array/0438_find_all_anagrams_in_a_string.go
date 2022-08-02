package array

func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	l, r, matched, res := 0, 0, 0, make([]int, 0)

	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	for r < len(s) {
		e := s[r]
		r++
		if need[e] > 0 {
			window[e]++
			if window[e] == need[e] {
				matched++
			}
		}

		if r-l == len(p) {
			if matched == len(need) {
				res = append(res, l)
			}

			d := s[l]
			l++
			if need[d] > 0 {
				if window[d] == need[d] {
					matched--
				}
				window[d]--
			}
		}
	}

	return res
}
