package binaryTree

func solveNQueens(n int) [][]string {
	res := make([][]string, n)
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		tmp := make([]byte, n)
		for j := 0; j < n; j++ {
			tmp[j] = '.'
		}
		board[i] = tmp
	}

	var valid = func(row, col int) bool {
		for i := 0; i < row; i++ {
			if board[row][col] == 'Q' {
				return false
			}
		}
		for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[row][col] == 'Q' {
				return false
			}
		}
		for i, j := row, col; i < n && j < n; i, j = i+1, j+1 {
			if board[row][col] == 'Q' {
				return false
			}
		}
		return true
	}
	var backTrack func(row int)
	backTrack = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i := 0; i < n; i++ {
				tmp = append(tmp, string(board[i]))
			}
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if !valid(row, i) {
				continue
			}

			board[row][i] = 'Q'
			backTrack(row + 1)
			board[row][i] = '.'
		}
	}

	backTrack(0)
	return res
}
