#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "../include/cge.h"

#define TEST 0

#if TEST
	#define FN "./test"
#else
	#define FN "./input"
#endif

int comp_long(const void* a, const void* b)
{
	return (*(long*)a) > (*(long*)b);
}

int mult_diffs(long* nums, int n)
{
	qsort(nums, n, sizeof(long), &comp_long);
	long* diffs = malloc(sizeof(long)*(n-1));

	int c1 = 1, c3 = 1;
	for (int i = 0; i < n-1; i++) {
		diffs[i] = nums[i+1] - nums[i];
		c1 += (diffs[i] == 1);
		c3 += (diffs[i] == 3);
	}

	free(diffs);
	printf("%d %d\n", c1, c3);
	return c1*c3;
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines = load_file(FN, 0, &n);
	long* nums = sstol(lines, n);

	int res = mult_diffs(nums, n);
	DUMP("%d", res);

	free(nums);
	ffree((void**)lines, n);
	return 0;
}
