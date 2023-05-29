package meeting

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	res := evalRPN(tokens)
	fmt.Println(res)
}
