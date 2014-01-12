package algoholic

// Merge Sort, O(n lg n) worst-case. Very beautiful.
func MergeSort(ns []int) []int {
	// Base case - an empty or length 1 slice is trivially sorted.
	if len(ns) < 2 {
		// We need not allocate memory here as the at most 1 element will only be referenced
		// once.
		return ns
	}

	half := len(ns) / 2

	// The wonder of merge sort - we sort each half of the slice using... merge sort :-)
	// Where is your God now?
	ns1 := MergeSort(ns[:half])
	ns2 := MergeSort(ns[half:])

	// We now have 2 separately sorted slices, merge them into one.
	return Merge(ns1, ns2)
}

// Merge, O(n), merges two sorted slices into one.
func Merge(ns1, ns2 []int) []int {
	length := len(ns1) + len(ns2)
	ret := make([]int, length)

	i, j := 0, 0

	// We iterate through each element of the returned slice, placing elements of each of the
	// input slices in their appropriate places.
	//
	// Loop Invariant: ret[:k] consists of ns1[:i] and ns2[:j] in sorted order.
	for k := 0; k < length; k++ {
		if j >= len(ns2) {
			copy(ret[k:], ns1[i:])
			break
		}
		if i >= len(ns1) {
			copy(ret[k:], ns2[j:])
			break
		}

		switch {
		case ns1[i] <= ns2[j]:
			ret[k] = ns1[i]
			i++
		default:
			ret[k] = ns2[j]
			j++
		}
	}

	// When the loop is complete, i == len(ns1), j == len(ns2). Therefore our loop invariant
	// determines that ret consists of ns1 and ns2 in sorted order, which matches the purpose
	// of the function.
	return ret
}
