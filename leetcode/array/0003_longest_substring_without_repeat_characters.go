package array

func lengthOfLongestSubstring(s string) int {
	max, l, r, window := 0, 0, 0, make(map[byte]int)
	for r < len(s) {
		e := s[r]
		if v, exi := window[e]; exi && v >= l {
			if size := r - l; size > max {
				max = size
			}
			l = v + 1
		}
		window[e] = r
		r++
	}

	if size := r - l; size > max {
		max = size
	}

	return max
}
