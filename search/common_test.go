package algoholic

import "testing"

type searchFunc func([]int, int) int

// Generate integers in range [from, to].
func generateRange(from, to int) []int {
	length := to - from + 1

	if length < 0 {
		panic("Invalid from, to.")
	}

	ret := make([]int, length)

	for i := 0; i < length; i++ {
		ret[i] = from + i
	}

	return ret
}

func correctlySearchesEmptySlice(t *testing.T, doSearch searchFunc) {
	if res := doSearch(nil, 123); res != -1 {
		t.Errorf("Search of nil returned %d, not expected -1.", res)
	}

	if res := doSearch([]int{}, 123); res != -1 {
		t.Errorf("Search of []int{} returned %d, not expected -1.", res)
	}
}

func correctlyFailsToFindMissingItem(t *testing.T, doSearch searchFunc) {
	for length := 0; length <= 100; length++ {
		ns := generateRange(1, length)
		if res := doSearch(ns, 0); res != -1 {
			t.Fatalf("Search of [1, %d] for 0 returned %d, not expected -1.",
				length, res)
		}
		if res := doSearch(ns, length+1); res != -1 {
			t.Fatalf("Search of [1, %d] for %d returned %d, not expected -1.",
				length, length+1, res)
		}
	}
}

func correctlyFindsItem(t *testing.T, doSearch searchFunc) {
	for length := 1; length <= 1e3; length++ {
		ns := generateRange(1, length)
		for i := 1; i <= length; i++ {
			if res := doSearch(ns, i); res != i-1 {
				t.Fatalf("Search of [1, %d] for %d returned %d, not expected %d.",
					length, i, res, i-1)
			}
		}
	}
}
