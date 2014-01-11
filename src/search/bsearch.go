package algoholic

// Binary Search - Given a sorted input slice and a sought value, return the index of that
// value in the slice or -1 if it can't be found.

// Binary Search, recursive implementation. O(lg n) worst-case.
func BinarySearchRecursive(ns []int, n int) int {
	// Declare our recursive function here so we can pass an offset parameter. This accounts
	// for the offset of the slice from the original input so we can return the correct index.
	var search func([]int, int, int) int
	search = func(ns []int, n, offset int) int {
		// If there's nothing to search, then we have failed to find the sought value.
		if len(ns) == 0 {
			return -1
		}

		// Look at the value in the middle of the slice.
		half := len(ns) / 2
		val := ns[half]

		switch {
		case n == val:
			// We have found the sought value.
			return half + offset
		case n < val:
			// Our sought value is less than the value at the halfway point, so we search the
			// slice below the halfway point.
			return search(ns[:half], n, offset)
		default: // n > val
			// Our sought value is greater than the value at the halfway point, so we search
			// the slice above the halfway point.
			return search(ns[half+1:], n, offset+half+1)
		}
	}

	// Kick off the search with inital offset of 0.
	return search(ns, n, 0)
}

// Binary Search, iterative implementation. O(lg n) worst-case.
func BinarySearchIterative(ns []int, n int) int {
	// Trivial case - The input slice is empty. Nothing to find.
	if len(ns) == 0 {
		return -1
	}

	// We track the range we are searching in. Initially this is the whole slice.
	from, to := 0, len(ns)-1

	for {
		// Avoid overflow.
		mid := to - (to-from)/2
		val := ns[mid]

		switch {
		case n == val:
			// We have found the sought value.
			return mid
		case n < val:
			// Our sought value is less than the value at the halfway point, so we set our
			// bounds to the slice below the halfway point.
			to = mid - 1
		default: // n > val
			// Our sought value is greater than the value at the halfway point, so we set our
			// bounds to the slice above the halfway point.
			from = mid + 1
		}

		// If the range is now invalid, the search has failed.
		if from > to {
			return -1
		}
	}
}
