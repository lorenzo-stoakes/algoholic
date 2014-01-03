package algoholic

import "testing"

// For convenience implement variation on isort which returns the slice once sorted.
func isort2(ns []int) []int {
	// In-place.
	InsertionSort(ns)
	return ns
}

func TestInsertionSortSortsReversedInts(t *testing.T) {
	correctlySortsReversedInts(t, 1e4, isort2)
}

func BenchmarkInsertionSort(b *testing.B) {
	benchmarkSortingReversedInts(b, isort2)
}
