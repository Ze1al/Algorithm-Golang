package offer

import (
	"fmt"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	res := minPathSum(grid)

	fmt.Println(res)
}
