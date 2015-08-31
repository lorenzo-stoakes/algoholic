#ifndef __sort_h
#define __sort_h

#include <stdlib.h>

typedef int *(*sort_fn_t)(int[], const ssize_t);

int *algoholic_isort(int ns[], const ssize_t len);

#endif
