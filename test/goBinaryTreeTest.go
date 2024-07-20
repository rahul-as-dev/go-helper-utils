package test

import (
	"github.com/rahul-as-dev/go-helper-utils/cmd/utils"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	tree := utils.NewBinaryTree(10)
	if tree == nil {
		t.Fatalf("Expected non-nil tree")
	}
	if tree.Value != 10 {
		t.Fatalf("Expected value 10, got %v", tree.Value)
	}
}

func TestSearch(t *testing.T) {
	// Test searching for a value in the binary tree
	tree := utils.NewBinaryTree(1)
	tree.Left = utils.NewBinaryTree(2)
	tree.Right = utils.NewBinaryTree(3)
	tree.Left.Left = utils.NewBinaryTree(4)
	tree.Left.Right = utils.NewBinaryTree(5)

	intCompare := func(a, b int) bool {
		return a == b
	}

	if !tree.SearchByComparator(5, intCompare) {
		t.Fatalf("Expected to find value 5")
	}
	if tree.SearchByComparator(6, intCompare) {
		t.Fatalf("Did not expect to find value 6")
	}
}
