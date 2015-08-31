package algoholic

import "testing"

func TestEmptyLinearSearch(t *testing.T) {
	correctlySearchesEmptySlice(t, LinearSearch)
}

func TestLinearSearchCantFindMissingItem(t *testing.T) {
	correctlyFailsToFindMissingItem(t, LinearSearch)
}

func TestLinearSearchFindsItem(t *testing.T) {
	correctlyFindsItem(t, LinearSearch)
}

func BenchmarkLinearSearch(b *testing.B) {
	benchmarkSearch(b, LinearSearch)
}
