#include <stdio.h>
#include <stdlib.h>

const int N = 1e4;

// Insertion Sort, O(n^2) worst-case. Interestingly, O(n) in the case of an already-sorted
// input.
void
isort(int *ns, int n)
{
	// The algorithm divides the slice into sorted and unsorted elements. The sorted portion
	// starts trivially with ns[0:1] and we iterate through the remaining unsorted elements
	// one-by-one positioning each into its correct position in the sorted portion.

    int i, j, key;

	// Loop Invariant: ns[0:i] is sorted.
    for(i = 1; i < n; i++) {
        key = ns[i];
		// We loop over each of the sorted elements which belong to the 'right' of the key
		// and move them up a place overwriting their neighbour.
        for(j = i-1; j >= 0 && key < ns[j]; j--)
            ns[j+1] = ns[j];
		// Finally we have the correct index to place our key value.
        ns[j+1] = key;
		// Now the loop invariant is established for ns[0:i+1] and thus we can increment i and
		// loop again.
    }
	// Once the loop is complete the loop invariant specifies that ns[0:n] is sorted, i.e. ns
	// is sorted and the sort has succeeded.
}

// For now, this code runs a simple test which sorts a reversed set of integers and ensures
// they are correctly sorted afterward.
int
main(void)
{
    int i;
    int *ns = malloc(sizeof(int)*N);
    for(i = 0; i < N; i++) {
        ns[i] = N-i;
    }

    isort(ns, N);

    for(i = 0; i < N; i++) {
        if(ns[i] != i+1) {
            fprintf(stderr, "Index %d is %d, expected %d.\n", i, ns[i], i+1);
            exit(EXIT_FAILURE);
        }
    }

    return EXIT_SUCCESS;
}
