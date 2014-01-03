package algoholic

import "testing"

func TestEmptyBinarySearchRecursive(t *testing.T) {
	correctlySearchesEmptySlice(t, BinarySearchRecursive)
}

func TestBinarySearchRecursiveCantFindMissingItem(t *testing.T) {
	correctlyFailsToFindMissingItem(t, BinarySearchRecursive)
}

func TestBinarySearchRecursiveFindsItem(t *testing.T) {
	correctlyFindsItem(t, BinarySearchRecursive)
}
