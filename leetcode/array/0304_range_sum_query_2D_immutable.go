package array

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix)+1, len(matrix[0])+1
	tmp := make([][]int, m)
	for i := 0; i < m; i++ {
		tmp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			tmp[i][j] = matrix[i-1][j-1] + tmp[i-1][j] + tmp[i][j-1] - tmp[i-1][j-1]
		}
	}

	return NumMatrix{preSum: tmp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}
