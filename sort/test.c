#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "sort.h"
#include "test.h"

/*
 * Sort a reversed set of integers and check that they are correctly sorted
 * afterwards.
 */
static void run_test(const sort_fn_t sort)
{
	const int len = 1e4;
	int *ns = gen_array(len, ARR_REVERSE_SORTED);

	ns = sort(ns, len);
	check_array(ns, len);
	free(ns);
}

int main(void)
{
	run_test(algoholic_isort);

	return EXIT_SUCCESS;
}
