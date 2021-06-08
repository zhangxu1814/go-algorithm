package reverse_integer

func reverse(x int) int {
	i := 0
	for x != 0 {
		i = i*10 + x%10
		x = x / 10
	}
	if i > 1<<31-1 || i < -(1<<31) {
		return 0
	}
	return i
}
