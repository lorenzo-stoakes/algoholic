package algoholic

import "testing"

const N = 1e4

// A simple test which sorts a reversed set of integers and ensures they are correctly sorted
// afterward.
func TestSorts(t *testing.T) {
	ns := make([]int, N)

	for i := 0; i < N; i++ {
		ns[i] = N - i
	}

	InsertionSort(ns)

	for i := 0; i < N; i++ {
		if ns[i] != i+1 {
			t.Errorf("Index %d is not sorted = %d, expected %d.", i, ns[i], i+1)
		}
	}
}
