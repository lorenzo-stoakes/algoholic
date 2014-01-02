package algoholic

import "testing"

func TestInsertionSortSortsReversedInts(t *testing.T) {
	correctlySortsReversedInts(t, 1e4, func(ns []int) []int {
		InsertionSort(ns)
		return ns
	})
}
