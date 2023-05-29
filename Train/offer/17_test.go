package offer

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	res := searchInsert(nums, 5)
	fmt.Println(res)
}
