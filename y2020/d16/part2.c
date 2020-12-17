#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

typedef struct __range range;
typedef struct __range {
	int l, h;
} range;

range* init_range(short l, short h)
{
	range* r = malloc(sizeof(range));
	r->l = l; r->h = h;
	return r;
}

int rangecmp(const void* a, const void* b)
{
	range* ra = (range*)a;
	range* rb = (range*)b;

	//printf("(%d, %d) < (%d, %d)\n", ra->l, ra->h, rb->l, rb->h);
	return ra->l > rb->l || (ra->l == rb->l && ra->h > rb->h);
}

#define ON_RANGE( r, x ) ((r.l <= x) && (x <= r.h))

int on_ranges(range* ranges, int n, int x)
{
	int res = 0;
	for (int i = 0; i < n; i++) {
		//printf("(%d, %d)\n", ranges[i].l, ranges[i].h);
		res |= ON_RANGE(ranges[i], x);
	}
	return res;
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

	char** toks; int m;
	long* nums;
	int i = 0, n_rules = 0, n_ranges = 0;
	int l1, h1, l2, h2;

	// Count rules
	while (i < n && strcmp(lines[i], "") != 0) {
		i++; n_rules++;
	}
	i = 0;
	n_ranges = n_rules*2;

	// Read rules
	range* ranges = calloc(1, sizeof(range)*n_ranges);
	while (i < n && strcmp(lines[i], "") != 0) {
		toks = split(lines[i], ":", &m);
		sscanf(toks[1], " %d-%d or %d-%d", &l1, &h1, &l2, &h2);
		FFREE(toks, m);
		ranges[i*2] = (range){ .l = l1, .h = h1 };
		ranges[(i*2)+1] = (range){ .l = l2, .h = h2 };
		i++;
	}
	i++;
	i++;

	// Read my ticket
	int myticket;
	while (i < n && strcmp(lines[i], "") != 0) {
		// NOP
		myticket = i;
		i++;
	}
	i++;
	i++;

	int ret = 0, sum = 0, valid = 1, k = 0, ioff = i;
	int* valid_inds = malloc(sizeof(int)*(n-ioff));
	long** valid_nums = calloc(1, sizeof(long*)*(n-ioff));
	// Read nearby tickets
	while (i < n && strcmp(lines[i], "") != 0) {
		toks = split(lines[i], ",", &m);
		nums = sstol(toks, m);
		valid = 1;
		for (int j = 0; j < m; j++) {
			if (!on_ranges(ranges, n_ranges, nums[j])) {
				sum += nums[j];
				valid = 0;
			}
		}

		if (valid) {
			valid_nums[k] = nums;
		} else {
			free(nums);
		}

		FFREE(toks, m);
		i++; k++;
	}
	FFREE(lines, n);

	for (int i = 0; i < k; i++) {
		for (int j = 0; j < m; j++) {
			printf("%ld ", valid_nums[i][j]);
		}
		printf("\n");
	}

	DUMP("%d", sum);

	free(ranges);
	return 0;
}
