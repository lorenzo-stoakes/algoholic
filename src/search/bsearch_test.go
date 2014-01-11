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

func BenchmarkBinarySearchRecursive(b *testing.B) {
	benchmarkSearch(b, BinarySearchRecursive)
}

func TestEmptyBinarySearchIterative(t *testing.T) {
	correctlySearchesEmptySlice(t, BinarySearchIterative)
}

func TestBinarySearchIterativeCantFindMissingItem(t *testing.T) {
	correctlyFailsToFindMissingItem(t, BinarySearchIterative)
}

func TestBinarySearchIterativeFindsItem(t *testing.T) {
	correctlyFindsItem(t, BinarySearchIterative)
}

func BenchmarkBinarySearchIterative(b *testing.B) {
	benchmarkSearch(b, BinarySearchIterative)
}
