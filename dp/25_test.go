package dp

import (
	"fmt"
	"testing"
)

func Test_maxProfit2(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit2(prices)
	fmt.Printf("res=%d", res)
}
