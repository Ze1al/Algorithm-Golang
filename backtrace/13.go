package backTrace

import "strings"

var res13 [][]string

// N皇后 https://leetcode.cn/problems/n-queens/
func solveNQueens(n int) [][]string {
	res13 = [][]string{}
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	backTracking13(0, n, board)
	return res13
}

// row第几行, path为缓存记录，board为棋盘
func backTracking13(row int, n int, board [][]string) {
	if row == n {
		temp := make([]string, n)
		for i := 0; i < len(board); i++ {
			temp[i] = strings.Join(board[i], "")
		}

		res13 = append(res13, temp)
		return
	}

	for i := 0; i < n; i++ {
		if isValidBoard(row, i, n, board) {
			board[row][i] = "Q"
			backTracking13(row+1, n, board)
			board[row][i] = "."
		}
	}
}

func isValidBoard(row, col, n int, board [][]string) bool {
	// 检查行
	for i := 0; i < n; i++ {
		if i != col && board[row][i] == "Q" {
			return false
		}
	}

	// 检查列
	for i := 0; i < n; i++ {
		if i != row && board[i][col] == "Q" {
			return false
		}
	}

	// 检查斜线, 45度
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	// 检查斜线, 135度
	for i, j := row-1, col+1; i >= 0 && j <= n-1; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
