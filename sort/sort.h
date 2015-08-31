#ifndef __sort_h
#define __sort_h

enum array_type { ARR_REVERSE_SORTED, ARR_SORTED, ARR_RANDOM };
typedef int *(*sort_fn_t)(int[], const int);

int *isort(int ns[], const int len);

#endif
