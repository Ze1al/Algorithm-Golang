package double_point

import (
	"fmt"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}

	res := sortedSquares(nums)
	fmt.Printf("res=%v", res)
}
