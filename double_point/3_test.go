package double_point

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

	fmt.Printf("res=%v", nums)
}
