package treenode

import (
	"fmt"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	res := binaryTreePaths(root)
	fmt.Printf("res=%v", res)
}
