#include <stdbool.h>

#include "sort.h"
#include "test.h"

int main(void)
{
	if (!test_sort(algoholic_isort, 1e4, ARR_REVERSE_SORTED))
		test_failure("Insertion Sort, Reverse Sorted Array", false);

	return EXIT_SUCCESS;
}
