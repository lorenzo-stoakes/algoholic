#ifndef __test_h
#define __test_h

#include <stdbool.h>
#include "sort.h"

bool check_array(const int ns[], const int len);
int next_val(int index, const int len, const enum array_type type);
int *gen_array(const int len, const enum array_type type);
bool test_sort(const sort_fn_t sort, const int len, const enum array_type type);

#endif
