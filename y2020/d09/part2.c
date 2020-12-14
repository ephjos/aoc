#include <limits.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

#define TEST 0

#if TEST == 1
#define N_PREAMBLE 5
#define FN "./test"
#else
#define N_PREAMBLE 25
#define FN "./input"
#endif


int can_sum_to(long long* nums, int n, int idx, long long t)
{
	long long first = 0;
	int lim = idx;
	for (int i = idx-N_PREAMBLE; i < lim; i++) {
		first = t-nums[i];
		for (int j = i+1; j < lim; j++) {
			if (first == nums[j]) return 1;
		}
	}
	return 0;
}

long long find_XMAS_weakness(long long* nums, int n)
{
	for (int i = N_PREAMBLE; i < n; i++) {
		if (!can_sum_to(nums, n, i, nums[i])) return nums[i];
	}
	return -1;
}

long long find_cont_to(long long* nums, int n, long long weakness)
{
	// 2D array of partial sums
	// d[i][j] represents the partial sum from index i to j
	long long** d = malloc(sizeof(long long*)*n);
	for (int i = 0; i < n; i++) {
		d[i] = calloc(1, sizeof(long long)*n);
	}

	int i, j;
	int done = 0;
	for (j = 1; j < n; j++) {
		for (i = 0; i < j; i++) {
			// Build all possible partial sums from current sums
			d[i][j] = d[i][j-1] + nums[j];
			if ((done = d[i][j] == weakness)) break;
		}
		if (done) break;
	}
	ffree((void**)d, n);

	long long min = LONG_MAX;
	long long max = LONG_MIN;

	for (int x = i; x <= j; x++) {
		min = MIN(min, nums[x]);
		max = MAX(max, nums[x]);
	}

	return min+max;
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines = load_file(FN, 0, &n);
	long long* nums = sstoll(lines, n);

	int weakness = find_XMAS_weakness(nums, n);
	int full_weakness = find_cont_to(nums, n, weakness);
	DUMP("%d", full_weakness);

	free(nums);
	ffree((void**)lines, n);
	return 0;
}
