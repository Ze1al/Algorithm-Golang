package dp

import (
	"fmt"
	"testing"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}

	res := uniquePathsWithObstacles(obstacleGrid)
	fmt.Printf("res=%v", res)
}
