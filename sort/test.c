#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "sort.h"
#include "test.h"

int main(void)
{
	test_sort(algoholic_isort, 1e4, ARR_REVERSE_SORTED);

	return EXIT_SUCCESS;
}
