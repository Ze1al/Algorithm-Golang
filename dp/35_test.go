package dp

import (
	"fmt"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	res := isSubsequence("abc", "ahbgdc")
	fmt.Printf("res=%v", res)
}
