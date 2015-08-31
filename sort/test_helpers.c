#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "test.h"

bool check_array(const int ns[], const size_t len)
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

int next_val(int index, const size_t len, const enum array_type type)
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

int *gen_array(const size_t len, const enum array_type type)
{
	size_t i;
	int *ret = calloc(sizeof(int), len);

	for (i = 0; i < len; i++)
		ret[i] = next_val(i, len, type);

	return ret;
}

bool test_sort(const sort_fn_t sort, const size_t len,
	const enum array_type type)
{
	bool ret;
	int *ns = gen_array(len, type);

	ns = sort(ns, len);
	ret = check_array(ns, len);
	free(ns);

	return ret;
}
