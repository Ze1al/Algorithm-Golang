package offer

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])
	total := row * col

	res := make([]int, total)

	direct := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	index := 0

	i, j := 0, 0

	for num := 0; num < total; num++ {
		res[num] = matrix[i][j]

		matrix[i][j] = -101

		nextI, nextJ := direct[index][0]+i, direct[index][1]+j
		if nextI < 0 || nextJ < 0 || nextI >= row || nextJ >= col || matrix[nextI][nextJ] == -101 {
			index = (index + 1) % 4
		}

		i = direct[index][0] + i
		j = direct[index][1] + j
	}
	return res
}
