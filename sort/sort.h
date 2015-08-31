#ifndef __sort_h
#define __sort_h

#include <stdlib.h>

enum array_type { ARR_REVERSE_SORTED, ARR_SORTED, ARR_RANDOM };
typedef int *(*sort_fn_t)(int[], const ssize_t);

int *algoholic_isort(int ns[], const ssize_t len);

#endif
