package algoholic

import "math/rand"

// In-place Fisher-Yates shuffle.
// O(n) assuming an O(1) PRNG.
func ShuffleFisherYates(ns []int) {
	for i := len(ns) - 1; i >= 1; i-- {
		j := rand.Int31n(int32(i + 1))
		ns[i], ns[j] = ns[j], ns[i]
	}
}
