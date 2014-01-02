package algoholic

// Merge Sort, O(n lg n) worst-case. Very beautiful.
func MergeSort(ns []int) []int {
	if len(ns) < 2 {
		return ns
	}

	half := len(ns) / 2

	ns1 := MergeSort(ns[:half])
	ns2 := MergeSort(ns[half:])

	return Merge(ns1, ns2)
}

func Merge(ns1, ns2 []int) []int {
	length := len(ns1) + len(ns2)
	ret := make([]int, length)

	i, j := 0, 0

	for k := 0; k < length; k++ {
		switch {
		case j >= len(ns2) || ns1[i] <= ns2[j]:
			ret[k] = ns1[i]
			i++
		case i >= len(ns1) || ns2[j] <= ns1[i]:
			ret[k] = ns2[j]
			j++
		}
	}

	return ret
}
