package offer

import (
	"fmt"
	"testing"
)

func Test_replaceWords(t *testing.T) {
	dictionary := []string{"a", "aa", "aaa", "aaaa"}
	sentence := "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"

	res := replaceWords(dictionary, sentence)
	fmt.Println(res)
}
