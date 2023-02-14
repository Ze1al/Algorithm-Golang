package treenode

import (
	"fmt"
	"testing"
)

func Test_findBottomLeftValue(t *testing.T) {
	root := &TreeNode{Val: 0}
	res := findBottomLeftValue(root)
	fmt.Printf("res=%v", res)
}
