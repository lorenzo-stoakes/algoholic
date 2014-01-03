package algoholic

import "testing"

func TestEmptyBinarySearch(t *testing.T) {
	correctlySearchesEmptySlice(t, BinarySearch)
}

func TestBinarySearchCantFindMissingItem(t *testing.T) {
	correctlyFailsToFindMissingItem(t, BinarySearch)
}

func TestBinarySearchFindsItem(t *testing.T) {
	correctlyFindsItem(t, BinarySearch)
}
