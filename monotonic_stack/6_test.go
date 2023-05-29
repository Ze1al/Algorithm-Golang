package monotonic_stack

import (
	"fmt"
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	height := []int{2, 4}
	res := largestRectangleArea(height)
	fmt.Printf("res=%d", res)
}
