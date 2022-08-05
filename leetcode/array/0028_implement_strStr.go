package array

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	l, r := 0, len(needle)
	for r <= len(haystack) {
		if haystack[l:r] == needle {
			return l
		}
		l++
		r++
	}
	return -1
}
