package monotonic_stack

import (
	"fmt"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	nums := []int{1, 2, 1}
	res := nextGreaterElements(nums)
	fmt.Println(res)
}
