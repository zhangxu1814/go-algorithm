package array

func reverseWords(s string) string {
	res := ""
	s = " " + s + " "
	l, r := len(s)-1, len(s)-1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == ' ' {
			r, l = l, i
			if v := s[l+1 : r+1]; v != " " {
				res += v
			}
		}
	}
	return res[:len(res)-1]
}
