package double_point

import (
	"fmt"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	ransomNote := "aa"
	magazine := "aab"

	res := canConstruct(ransomNote, magazine)
	fmt.Printf("res=%v", res)
}
