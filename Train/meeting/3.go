package meeting

// 猿辅导
// 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])

	i, j := 0, 0 // 起点

	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[0]))
	}
	// 方向
	d := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	dIndex := 0
	total := row * col
	res := make([]int, total)

	for t := 0; t < total; t++ {
		res[t] = matrix[i][j]
		visited[i][j] = true

		// 下一个节点的地址
		nexti, nextj := i+d[dIndex][0], j+d[dIndex][1]

		// 转变方向
		if nexti >= row || nextj >= col || nexti < 0 || nextj < 0 || visited[nexti][nextj] == true {
			dIndex = (dIndex + 1) % 4
		}

		i = i + d[dIndex][0]
		j = j + d[dIndex][1]
	}
	return res
}
