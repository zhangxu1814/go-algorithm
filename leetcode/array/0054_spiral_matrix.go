package array

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	up, down, left, right := 0, m-1, 0, n-1
	res, cur := make([]int, m*n), 0
	for cur < m*n {
		if up <= down {
			for i := left; i <= right; i++ {
				res[cur] = matrix[up][i]
				cur++
			}
			up++
		}
		if left <= right {
			for i := up; i <= down; i++ {
				res[cur] = matrix[i][right]
				cur++
			}
			right--
		}
		if up <= down {
			for i := right; i >= left; i-- {
				res[cur] = matrix[down][i]
				cur++
			}
			down--
		}
		if left <= right {
			for i := down; i >= up; i-- {
				res[cur] = matrix[i][left]
				cur++
			}
			left++
		}
	}
	return res
}
