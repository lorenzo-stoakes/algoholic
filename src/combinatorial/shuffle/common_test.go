package algoholic

import (
	"math"
	"testing"
)

// Generate a slice [1, n].
func generateSlice(n int) []int {
	ret := make([]int, n)

	for i := 0; i < n; i++ {
		ret[i] = i + 1
	}

	return ret
}

// Ensure that the specified shuffle function distributes the specified length list over the
// specified number of iterations with specified maximum error in
func checkShufflesEvenly(t *testing.T, shuffle func([]int), length, iters int, maxErr float64) {
	// Create slice equal to [1, length].
	ns := generateSlice(length)

	// For each index in [0, length) count occurrences of each number in the range [1, length].
	counts := make([][]int, length)

	for i := 0; i < length; i++ {
		counts[i] = make([]int, length)
	}

	for i := 0; i < iters; i++ {
		shuffle(ns)

		for j, n := range ns {
			counts[n-1][j]++
		}
	}

	for i, countSet := range counts {
		for j, count := range countSet {
			// Actual probability of j+1 occurring in position i.
			actual := float64(count) / float64(iters)
			// Expect an even distribution.
			expected := 1 / float64(length)
			// Percentage variation from expectation.
			err := math.Abs((actual - expected) / expected)

			if err > maxErr {
				t.Errorf("Occurrence of %d at index %d varies from expectation by %0.2f%%, not expected max error of %0.2f%%.",
					j+1, i, err*100, maxErr*100)
			}
		}
	}
}
