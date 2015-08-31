#ifndef __test_h
#define __test_h

#include <stdbool.h>
#include "sort.h"

bool check_array(const int ns[], const size_t len);
int next_val(int index, const size_t len, const enum array_type type);
int *gen_array(const size_t len, const enum array_type type);
bool test_sort(const sort_fn_t sort, const size_t len,
	const enum array_type type);

#endif
