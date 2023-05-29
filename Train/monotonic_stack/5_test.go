package monotonic_stack

import (
	"fmt"
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	nums := []int{2, 1, 5, 6, 2, 3}
	res := largestRectangleArea(nums)
	fmt.Println(res)
}
