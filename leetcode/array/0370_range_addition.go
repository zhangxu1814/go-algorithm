package array

func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length)

	var start, end, value int
	for i := 0; i < len(updates); i++ {
		start, end, value = updates[i][0], updates[i][1], updates[i][2]
		diff[start] += value
		if end+1 < length {
			diff[end+1] -= value
		}
	}

	for i := 1; i < length; i++ {
		diff[i] += diff[i-1]
	}

	return diff
}
