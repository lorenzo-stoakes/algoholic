package algoholic

import "testing"

// A simple test which sorts a reversed set of integers and ensures they are correctly sorted
// afterward.
func TestMergeSortSortsReversedInts(t *testing.T) {
	correctlySortsReversedInts(t, 1e6, MergeSort)
}
