package double_point

import (
	"fmt"
	"testing"
)

func Test_isHappy(t *testing.T) {
	n := 19
	res := isHappy(n)
	fmt.Printf("res=%v", res)
}
