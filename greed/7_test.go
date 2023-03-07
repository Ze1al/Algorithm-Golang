package greed

import (
	"fmt"
	"testing"
)

func Test_largestSumAfterKNegations(t *testing.T) {
	num := []int{2, -3, -1, 5, -4}
	res := largestSumAfterKNegations(num, 2)
	fmt.Printf("res=%v", res)
}
