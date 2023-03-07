package greed

import (
	"fmt"
	"testing"
)

func Test_candy(t *testing.T) {
	ratings := []int{1, 2, 2, 5, 4, 3, 2}
	res := candy(ratings)
	fmt.Printf("res=%v", res)
}
