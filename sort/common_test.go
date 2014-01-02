package algoholic

import "testing"

// A simple test which sorts a reversed set of integers and ensures they are correctly sorted
// afterward.
func correctlySortsReversedInts(t *testing.T, max int, doSort func([]int) []int) {
	ns := make([]int, max)

	for i := 0; i < max; i++ {
		ns[i] = max - i
	}

	sorted := doSort(ns)

	for i := 0; i < max; i++ {
		if sorted[i] != i+1 {
			t.Fatalf("Index %d is not sorted = %d, expected %d.", i, sorted[i], i+1)
		}
	}
}
