package offer

import (
	"fmt"
	"testing"
)

func Test_minimumTotal(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	res := minimumTotal(triangle)
	fmt.Println(res)
}
