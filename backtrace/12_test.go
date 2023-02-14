package backTrace

import (
	"fmt"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}

	res := permuteUnique(nums)
	fmt.Printf("res=%v", res)
}
