package monotonic_stack

import (
	"fmt"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	num := []int{1, 2, 1}
	res := nextGreaterElements(num)
	fmt.Printf("res=%v", res)
}
