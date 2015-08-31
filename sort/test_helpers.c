#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#include "test.h"

static bool check_array(const int ns[], const size_t len)
{
	size_t i;
	bool ret = true;

	for (i = 1; i < len; i++) {
		if (ns[i - 1] > ns[i]) {
			fprintf(stderr, "ns[%zu], ns[%zu] == %d, %d\n", i,
				i + 1, ns[i], ns[i + 1]);
			ret = false;
		}
	}

	return ret;
}

static int next_val(int index, const size_t len, const enum array_type type)
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

static int *gen_array(const size_t len, const enum array_type type)
{
	size_t i;
	int *ret = calloc(sizeof(int), len);

	for (i = 0; i < len; i++)
		ret[i] = next_val(i, len, type);

	return ret;
}

static bool test_sort_timed(const sort_fn_t sort, const size_t len,
			const enum array_type type, bool check, time_t *takenp)
{
	bool ret = true;
	clock_t start, clocks;
	int *ns = gen_array(len, type);

	start = clock();
	ns = sort(ns, len);
	clocks = clock() - start;
	*takenp = 1000 * clocks / CLOCKS_PER_SEC;

	if (check)
		ret = check_array(ns, len);

	free(ns);
	return ret;
}

bool test_sort(const sort_fn_t sort, const size_t len,
	const enum array_type type)
{
	time_t dummy;

	return test_sort_timed(sort, len, type, true, &dummy);
}

void test_failure(const char *name, const bool fatal)
{
	fprintf(stderr, "\t- %s FAILED.\n", name);

	if (fatal)
		exit(EXIT_FAILURE);
}

size_t get_reasonable_length(sort_fn_t sort)
{
	size_t ret = 1000;
	time_t duration = 0;

	while (duration < MAX_SORT_DURATION_MS) {
		test_sort_timed(sort, ret, ARR_REVERSE_SORTED, false,
				&duration);

		if (duration < MAX_SORT_DURATION_MS / 2)
			ret *= 2;
		else
			ret *= 1.2;
	}

	return ret;
}
