package algoholic

import "testing"

func TestEmptyBinarySearch(t *testing.T) {
	if res := BinarySearch(nil, 123); res != -1 {
		t.Errorf("Binary search of nil resulted in %d, not expected -1.", res)
	}

	if res := BinarySearch([]int{}, 123); res != -1 {
		t.Errorf("Binary search of []int{} resulted in %d, not expected -1.", res)
	}
}

func TestBinarySearchCantFindMissingItem(t *testing.T) {
	for length := 0; length <= 100; length++ {
		ns := generateRange(1, length)
		if res := BinarySearch(ns, 0); res != -1 {
			t.Fatalf("Binary search of slice of [1, %d] for 0 returned %d, not expected -1.",
				length, res)
		}
		if res := BinarySearch(ns, length+1); res != -1 {
			t.Fatalf("Binary search of slice of [1, %d] for %d returned %d, not expected -1.",
				length, length+1, res)
		}
	}
}

func TestBinarySearchFindsItem(t *testing.T) {
	for length := 1; length <= 1e3; length++ {
		ns := generateRange(1, length)
		for i := 1; i <= length; i++ {
			if res := BinarySearch(ns, i); res != i-1 {
				t.Fatalf("Binary search did not find %d at index %d of [1, %d], returned %d.",
					i, i-1, length, res)
			}
		}
	}
}
