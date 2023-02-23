package dp

import (
	"fmt"
	"testing"
)

func Test_maxProfit3(t *testing.T) {
	price := []int{3, 3, 5, 0, 0, 3, 1, 4}
	res := maxProfit3(price)
	fmt.Printf("res=%d", res)
}
