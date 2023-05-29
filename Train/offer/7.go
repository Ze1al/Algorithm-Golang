package offer

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfs7(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func dfs7(grid [][]int, i, j int) int {
	var area int
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return area
	}

	// 代表从陆地走到海洋
	if grid[i][j] == 0 {
		return area
	}

	area += 1
	grid[i][j] = 0

	area += dfs7(grid, i+1, j)
	area += dfs7(grid, i-1, j)
	area += dfs7(grid, i, j+1)
	area += dfs7(grid, i, j-1)
	return area
}
