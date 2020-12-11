#include <math.h>
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

unsigned long tribonacci(int n)
{
	if (n == 0) return 0;
	if (n == 1) return 1;
	if (n == 2) return 1;
	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3);
}

long count_ways(long* nums, int n)
{
	long* diff = malloc(sizeof(long)*(n+1));
	for (int i = 0; i < n+1; i++) {
		diff[i] = nums[i+1] - nums[i];
	}

	unsigned long s = 1;
	int a = -1;
	for (int i = 0; i < n+1; i++) {
		if (diff[i] == 3) {
			if (i-a != 1) {
				s *= tribonacci(i-a);
			}
			a = i;
		}
	}

	free(diff);
	return s;
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines = load_file(FN, 0, &n);
	long* numsff = sstol(lines, n);
	qsort(numsff, n, sizeof(long), &comp_long);
	long* nums = malloc(sizeof(long)*(n+2));
	nums[0] = 0;
	for (int i = 1; i < n+1; i++) {
		nums[i] = numsff[i-1];
	}
	nums[n+1] = nums[n]+3;

	for (int i = 0; i < n+2; i++) {
		DUMP("%5ld", nums[i]);
	}

	unsigned long res = count_ways(nums, n);
	DUMP("%lu", res);

	free(nums);
	free(numsff);
	ffree((void**)lines, n);
	return 0;
}
