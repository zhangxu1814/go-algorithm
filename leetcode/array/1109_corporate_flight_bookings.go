package array

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)

	for i := 0; i < len(bookings); i++ {
		start, end, value := bookings[i][0]-1, bookings[i][1]-1, bookings[i][2]
		diff[start] += value
		if end+1 < n {
			diff[end+1] -= value
		}
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	return diff
}
