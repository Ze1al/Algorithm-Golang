package offer

import (
	"fmt"
	"testing"
)

func TestConstructor11(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")
	res := trie.Search("apple")
	fmt.Println(res)
}
