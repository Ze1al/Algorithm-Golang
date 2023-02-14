package treenode

import (
	"testing"
)

func Test_buildTree2(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	buildTree2(preorder, inorder)
}
