package stack

import (
	"fmt"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	s := []string{"2", "1", "+", "3", "*"}
	res := evalRPN(s)
	fmt.Printf("res=%d", res)
}
