package double_point

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	res := removeDuplicates(nums)
	fmt.Printf("res=%v", res)
}
