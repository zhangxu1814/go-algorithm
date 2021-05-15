package roman_to_integer

type stack struct {
	s   []byte
	len int
}

func constructor(s []byte) stack {
	return stack{
		s:   s,
		len: len(s),
	}
}

func (this *stack) pop() byte {
	res := this.top()
	this.s = this.s[:this.len-1]
	this.len--
	return res
}

func (this *stack) top() byte {
	if this.empty() {
		return ' '
	}
	return this.s[this.len-1]
}

func (this *stack) empty() bool {
	return this.len == 0
}

func romanToInt(s string) int {
	ss := constructor([]byte(s))
	res := 0

	for !ss.empty() {
		switch ss.pop() {
		case 'I':
			res += 1
		case 'V':
			switch ss.top() {
			case 'I':
				ss.pop()
				res += 4
			default:
				res += 5
			}
		case 'X':
			switch ss.top() {
			case 'I':
				ss.pop()
				res += 9
			default:
				res += 10
			}
		case 'L':
			switch ss.top() {
			case 'X':
				ss.pop()
				res += 40
			default:
				res += 50
			}
		case 'C':
			switch ss.top() {
			case 'X':
				ss.pop()
				res += 90
			default:
				res += 100
			}
		case 'D':
			switch ss.top() {
			case 'C':
				ss.pop()
				res += 400
			default:
				res += 500
			}
		case 'M':
			switch ss.top() {
			case 'C':
				ss.pop()
				res += 900
			default:
				res += 1000
			}
		}
	}

	return res
}
