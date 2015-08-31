#include "sort.h"

/*
 * Insertion Sort
 *
 * https://en.wikipedia.org/wiki/Insertion_sort
 *
 * Time
 *     Worst: O(n^2)
 *   Average: O(n^2)
 *      Best: O(n)
 *
 * Memory
 *     Worst: O(1)
 *   Average: O(1)
 *      Best: O(1)
 */
int *isort(int ns[], int n)
{
	int i, j, key;

	/*
	 * The key concept of insertion sort is to divide the input array into
	 * into sorted and unsorted subarrays and iteratively insert elements
	 * from the unsorted subarray into the sorted one.
	 *
	 * The subarray consisting of only the first element is trivially
	 * sorted so we begin by partitioning the input into one subarray of
	 * the first element and another of all remaining elements.
	 *
	 * By inserting each element of the unsorted subarray into its correct
	 * position in the sorted subarray we gradually grow the former and
	 * shrink the latter until we are left with sorted elements only - at
	 * this point the array is sorted and the algorithm has succeeded.
	 *
	 * We consider each element in the unsorted subarray from left to right
	 * (we term the unsorted element under consideration the 'key' value)
	 * against each element of the sorted subarray from right to left,
	 * moving the sorted subarray up a position on each occasion we find
	 * that the key is less than a sorted element and finally insert the key
	 * into this position.
	 *
	 * There are some interesting consequences of the design of insertion
	 * sort:
	 *
	 * 1. If an array is already sorted then we need only compare each
	 * unsorted element once against the last element of the sorted
	 * subarray, meaning we only have to perform n comparisons and the
	 * algorithm will have an O(n) running time.
	 *
	 * 2. We can avoid expensive swaps by simply storing the key element in
	 * a temporary variable and shifting elements up a place by overwriting
	 * the each element moving right to left starting with the key element
	 * for as long as the key is less than each element before finally
	 * placing the key element in the remaining position.
	 *
	 * 3. We can do all of this 'in place' by merely swapping around
	 * elements so the algorithm uses constant time memory in all cases.
	 */

	/* Loop Invariant: ns[0, i) is sorted. */
	for (i = 1; i < n; i++) {
		key = ns[i];
		/*
		 * We loop over each of the sorted elements which belong to the
		 * 'right' of the key and move them up a place overwriting their
		 * neighbour.
		 */
		for (j = i - 1; j >= 0 && key < ns[j]; j--)
			ns[j + 1] = ns[j];
		/* Finally we have the correct index to place our key value. */
		ns[j + 1] = key;
		/*
		 * Now the loop invariant is established for ns[0, i+1) and thus
		 * we can increment i and loop again.
		 */
	}

	/*
	 * Once the loop is complete the loop invariant specifies that ns[0, n)
	 * is sorted, i.e. ns is sorted and the sort has succeeded.
	 */
	return ns;
}
