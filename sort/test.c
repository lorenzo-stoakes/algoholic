#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "sort.h"

const int N = 1e4;

static bool check_array(const int ns[], const int len)
{
	int i;
	bool ret = true;

	for (i = 1; i < len; i++) {
		if (ns[i - 1] > ns[i]) {
			fprintf(stderr, "ns[%d], ns[%d] == %d, %d\n", i,
				i + 1, ns[i], ns[i + 1]);
			ret = false;
		}
	}

	return ret;
}

/*
 * Sort a reversed set of integers and check that they are correctly sorted
 * afterwards.
 */
static void run_test(const sort_fn_t sort)
{
	int i;
	int *ns = calloc(sizeof(int), N);

	for (i = 0; i < N; i++)
		ns[i] = N - i;

	ns = sort(ns, N);

	check_array(ns, N);

	free(ns);
}

int main(void)
{
	run_test(isort);

	return EXIT_SUCCESS;
}
