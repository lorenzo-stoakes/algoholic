#include <stdbool.h>

#include "sort.h"
#include "test.h"

int main(void)
{
	size_t len;

	len = get_reasonable_length(algoholic_isort);
	if (!test_sort(algoholic_isort, len, ARR_REVERSE_SORTED))
		test_failure("Insertion Sort, Reverse Sorted Array", false);

	return EXIT_SUCCESS;
}
