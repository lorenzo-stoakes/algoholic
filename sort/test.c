#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "sort.h"

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

static int next_val(int index, const int len, const enum array_type type)
{
	switch (type) {
	case ARR_REVERSE_SORTED:
		return len - index;
	case ARR_SORTED:
		return index + 1;
	case ARR_RANDOM:
		return rand();
	}

	fprintf(stderr, "Unknown array type %d, aborting!\n", type);
	exit(EXIT_FAILURE);
}

static int *gen_array(const int len, const enum array_type type)
{
	int i;
	int *ret = calloc(sizeof(int), len);

	for (i = 0; i < len; i++)
		ret[i] = next_val(i, len, type);

	return ret;
}

/*
 * Sort a reversed set of integers and check that they are correctly sorted
 * afterwards.
 */
static void run_test(const sort_fn_t sort)
{
	const int len = 1e4;
	int *ns = gen_array(len, ARR_REVERSE_SORTED);

	ns = sort(ns, len);
	check_array(ns, len);
	free(ns);
}

int main(void)
{
	run_test(isort);

	return EXIT_SUCCESS;
}
