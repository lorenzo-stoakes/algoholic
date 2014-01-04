package algoholic

import "testing"

type sortFunc func([]int) []int

const BENCHMARK_LENGTH = 1e4

func benchmarkSortingReversedInts(b *testing.B, doSort sortFunc) {
	slices := make([][]int, b.N)

	for i := 0; i < b.N; i++ {
		slices[i] = createReversedInts(BENCHMARK_LENGTH)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		doSort(slices[i])
	}
}

// A simple test which sorts a reversed set of integers and ensures they are correctly sorted
// afterward.
func correctlySortsReversedInts(t *testing.T, max int, doSort sortFunc) {
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
