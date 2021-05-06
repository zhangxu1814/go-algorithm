package backspace_string_compare

func backspaceCompare(s string, t string) bool {
	ss := []byte{}
	st := []byte{}
	for _, v := range []byte(s) {
		if v == '#' {
			if len(ss) > 0 {
				ss = ss[:len(ss)-1]
			}
			continue
		}
		ss = append(ss, v)
	}
	for _, v := range []byte(t) {
		if v == '#' {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
			continue
		}
		st = append(st, v)
	}

	return string(ss) == string(st)
}
