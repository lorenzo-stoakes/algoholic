#ifndef __sort_h
#define __sort_h

typedef int *(*sort_fn_t)(int[], const int);

int *isort(int ns[], const int n);

#endif
