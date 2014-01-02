package algoholic

import "testing"

func TestInsertionSortSortsReversedInts(t *testing.T) {
	correctlySortsReversedInts(t, 1e4, func(ns []int) []int {
		// In-place sort.
		InsertionSort(ns)
		return ns
	})
}
