package array

func longestPalindrome(s string) string {
	var max string
	for i := 0; i < len(s); i++ {
		tmp := find(s, i, i)
		b := find(s, i, i+1)
		if len(tmp) < len(b) {
			tmp = b
		}
		if len(tmp) > len(max) {
			max = tmp
		}
	}
	return max
}

func find(s string, l, r int) string {
	for l >= 0 && r <= len(s)-1 {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	l++
	return s[l:r]
}
