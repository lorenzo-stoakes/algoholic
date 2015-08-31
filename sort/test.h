#ifndef __test_h
#define __test_h

#include <stdbool.h>
#include "sort.h"

#define MAX_SORT_DURATION_MS 100

enum array_type { ARR_REVERSE_SORTED, ARR_SORTED, ARR_RANDOM };

bool test_sort(const sort_fn_t sort, const size_t len,
	const enum array_type type);
void test_failure(const char *name, const bool fatal);
size_t get_reasonable_length(sort_fn_t sort);

#endif
