#ifndef __test_h
#define __test_h

#include <stdbool.h>
#include "sort.h"

bool test_sort(const sort_fn_t sort, const size_t len,
	const enum array_type type);

#endif
