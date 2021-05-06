package decode_xored_array

func decode(encoded []int, first int) []int {
	result := append([]int{first}, encoded...)
	for i := 1; i < len(result); i++ {
		result[i] = result[i] ^ result[i-1]
	}
	return result
}
