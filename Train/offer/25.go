package offer

// 螺旋矩阵
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	direct := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	index := 0

	i, j := 0, 0
	for nums := 1; nums <= n*n; nums++ {
		matrix[i][j] = nums

		nextI, nextJ := direct[index][0]+i, direct[index][1]+j
		if nextI < 0 || nextJ < 0 || nextI >= n || nextJ >= n || matrix[nextI][nextJ] != 0 {
			index = (index + 1) % 4
		}

		i = direct[index][0] + i
		j = direct[index][1] + j
	}
	return matrix
}
