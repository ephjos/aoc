#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"
#include "cht.h"

// play a round
//
// @param nums: input numbers | takes ownership
// @param n: length of input
// @param nth: number of sequence to print
void playgame(long* nums, long n, long nth)
{
	cht* d = NEW_CHT();

	int i; long curr, prev, ls, li, t, x;

	// Starting numbers
	for (i = 0; i < n; i++){
		INS_CHT_I(d, nums[i], i);
		prev = nums[i];
		//DUMP("%ld", prev);
	}

	// Play til nth
	for (; i < nth; i++){
		if ((t = GET_CHT_I(d, prev)) == NOT_FOUND) {
			INS_CHT_I(d, prev, i-1);
			t = i-1;
		}
		x = i-t-1;
		INS_CHT_I(d, prev, i-1);
		prev = x;
		//DUMP("%ld", prev);
	}

	printf("%ldth number: %ld\n", nth, prev);

	free_cht(d);
	free(nums);
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines;

	if (argc == 3 &&
			(strcmp(argv[1], "-f") == 0 ||
			 strcmp(argv[1], "--file") == 0)) {
		lines = load_file(argv[2], 0, &n);
	} else {
		lines = load_file("./input", 0, &n);
	}

	int m;
	char** toks;
	for (int i = 0; i < n; i++) {
		toks = split(lines[i], ",", &m);
		playgame(sstol(toks, m), m, 2020);
		FFREE(toks, m);
	}

	FFREE(lines, n);
	return 0;
}
