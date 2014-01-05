package algoholic

// Linear Search - Given an input slice and a sought value, return the index of that value in
// the slice or -1 if it can't be found.

// Linear Search, O(n) worst-case.
func LinearSearch(ns []int, n int) int {
	// Simply iterate through every item in the slice, short-circuiting and returning should we
	// find the saught value.
	for i, m := range ns {
		if m == n {
			return i
		}
	}

	// We haven't found the value.
	return -1
}
