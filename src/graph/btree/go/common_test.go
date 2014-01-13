package algoholic

import (
	"math/rand"
	"testing"
)

func generateRandomTree(max int) *BinaryNode {
	ns := rand.Perm(max)
	return NewBinarySearchTree(ns)
}

// O(n^2)!
func assertSearchProperty(t *testing.T, node *BinaryNode) {
	if node == nil {
		return
	}

	node.Left.Walk(func(n int) {
		if n > node.Val {
			t.Fatalf("Tree fails search property, %d is in left subtree of %d.", n, node.Val)
		}
	})

	node.Right.Walk(func(n int) {
		if n <= node.Val {
			t.Fatalf("Tree fails search property, %d is in right subtree of %d.", n, node.Val)
		}
	})

	assertSearchProperty(t, node.Left)
	assertSearchProperty(t, node.Right)
}
