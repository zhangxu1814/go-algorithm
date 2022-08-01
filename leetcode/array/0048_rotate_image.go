package array

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n/2; i++ {
		m := n - 1 - i
		for j := 0; j < n; j++ {
			matrix[j][i], matrix[j][m] = matrix[j][m], matrix[j][i]
		}
	}
}
