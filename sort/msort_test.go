package algoholic

import "testing"

// A simple test which sorts a reversed set of integers and ensures they are correctly sorted
// afterward.
func TestMergeSortSortsReversedInts(t *testing.T) {
	correctlySortsReversedInts(t, 1e6, MergeSort)
}

func TestMergeMergesEmptySlicesCorrectly(t *testing.T) {
	const N = 10

	empty := []int{}
	ns := make([]int, N)

	for i := 0; i < N; i++ {
		ns[i] = i + 1
	}

	if zero := len(Merge(empty, empty)); zero != 0 {
		t.Errorf("Merge([]int{}, []int{}) has %d length, not expected 0.", zero)
	}

	merged := Merge(empty, ns)

	if !slicesAreEqual(ns, merged) {
		t.Error("Merge([]int{}, ns) does not return ns.")
	}

	merged = Merge(ns, empty)

	if !slicesAreEqual(ns, merged) {
		t.Error("Merge(ns, []int{}) does not return ns.")
	}
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkSortingReversedInts(b, MergeSort)
}
