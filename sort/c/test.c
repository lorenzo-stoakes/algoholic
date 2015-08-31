#include <stdio.h>
#include <stdlib.h>

#include "sort.h"

const int N = 1e4;

/*
 * Sort a reversed set of integers and check that they are correctly sorted
 * afterwards.
 */
int main(void)
{
	int i;
	int *ns = calloc(sizeof(int), N);

	for(i = 0; i < N; i++)
		ns[i] = N - i;

	isort(ns, N);

	for(i = 0; i < N; i++) {
		if(ns[i] != i+1) {
			fprintf(stderr, "Index %d is %d, expected %d.\n", i, ns[i], i+1);
			exit(EXIT_FAILURE);
		}
	}

	return EXIT_SUCCESS;
}
