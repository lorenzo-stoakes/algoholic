#include <stdio.h>
#include <stdlib.h>

const int N = 1e4;

void
isort(int *ns, int n)
{
    int i, j, key;

    for(i = 1; i < n; i++) {
        key = ns[i];
        for(j = i-1; j >= 0 && key < ns[j]; j--)
            ns[j+1] = ns[j];
        ns[j+1] = key;
    }
}

int
main(void)
{
    int i;
    int *ns = malloc(sizeof(int)*N);
    for(i = 0; i < N; i++) {
        ns[i] = N-i;
    }

    isort(ns, N);

    for(i = 0; i < N; i++) {
        if(ns[i] != i+1) {
            fprintf(stderr, "Index %d is %d, expected %d.\n", i, ns[i], i+1);
            exit(EXIT_FAILURE);
        }
    }

    return EXIT_SUCCESS;
}
