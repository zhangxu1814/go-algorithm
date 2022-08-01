package array

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	up, down, left, right, cur := 0, n-1, 0, n-1, 0
	for cur < n*n {
		if up <= down {
			for i := left; i <= right; i++ {
				matrix[up][i] = cur + 1
				cur++
			}
			up++
		}
		if left <= right {
			for i := up; i <= down; i++ {
				matrix[i][right] = cur + 1
				cur++
			}
			right--
		}
		if up <= down {
			for i := right; i >= left; i-- {
				matrix[down][i] = cur + 1
				cur++
			}
			down--
		}
		if left <= right {
			for i := down; i >= up; i-- {
				matrix[i][left] = cur + 1
				cur++
			}
			left++
		}
	}
	return matrix
}
