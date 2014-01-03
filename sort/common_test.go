package algoholic

import "testing"

func benchmarkSortingReversedInts(b *testing.B, doSort func([]int) []int) {
	ns := createReversedInts(b.N)
	b.ResetTimer()
	doSort(ns)
}

// A simple test which sorts a reversed set of integers and ensures they are correctly sorted
// afterward.
func correctlySortsReversedInts(t *testing.T, max int, doSort func([]int) []int) {
	ns := createReversedInts(max)

	sorted := doSort(ns)

	for i := 0; i < max; i++ {
		if sorted[i] != i+1 {
			t.Fatalf("Unsorted at index %d: got %d, expected %d.", i, sorted[i], i+1)
		}
	}
}

func createReversedInts(max int) []int {
	ret := make([]int, max)

	for i := 0; i < max; i++ {
		ret[i] = max - i
	}

	return ret
}

func slicesAreEqual(ns1, ns2 []int) bool {
	if ns1 == nil || ns2 == nil || len(ns1) != len(ns2) {
		return false
	}

	for i, n := range ns1 {
		if ns2[i] != n {
			return false
		}
	}

	return true
}
