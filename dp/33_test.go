package dp

import (
	"fmt"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"

	res := longestCommonSubsequence(text1, text2)
	fmt.Printf("res=%d", res)
}
