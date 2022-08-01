package array

func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right, matched := 0, 0, 0
	for right < len(s2) {
		e := s2[right]
		right++
		if _, exi := need[e]; exi {
			window[e]++
			if window[e] == need[e] {
				matched++
			}
		}

		if right-left == len(s1) {
			if matched == len(need) {
				return true
			}
			d := s2[left]
			left++
			if _, exi := need[d]; exi {
				if window[d] == need[d] {
					matched--
				}
				window[d]--
			}
		}
	}
	return false
}
