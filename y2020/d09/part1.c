#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

#define N_PREAMBLE 25

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
		//DUMP("%lld", nums[i]);
		if (!can_sum_to(nums, n, i, nums[i])) return nums[i];
	}
	return -1;
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines = load_file("./input", 0, &n);
	long long* nums = sstoll(lines, n);

	int weakness = find_XMAS_weakness(nums, n);
	DUMP("%d", weakness);

	ffree((void**)lines, n);
	return 0;
}
