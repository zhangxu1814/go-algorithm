package array

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)

	for i := 0; i < len(trips); i++ {
		value, start, end := trips[i][0], trips[i][1], trips[i][2]-1
		diff[start] += value
		if end+1 < 1001 {
			diff[end+1] -= value
		}
	}

	if diff[0] > capacity {
		return false
	}

	for i := 1; i < 1001; i++ {
		diff[i] += diff[i-1]
		if diff[i] > capacity {
			return false
		}
	}

	return true
}
