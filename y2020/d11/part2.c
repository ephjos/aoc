#include <assert.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

#include "../include/cge.h"

#define TEST 0
#if TEST
	#define FN "./test"
#else
	#define FN "./input"
#endif

#define EMPTY 'L'
#define OCC '#'
#define FLOOR '.'

void print_state(char** s, int n)
{
	//system("clear");
	for (int i = 0; i < n; i++) {
		printf("%s\n", s[i]);
	}
	printf("\n");
}

int states_equal(char** a, char** b, int n)
{
	if ((a == NULL && b != NULL) || (a != NULL && b == NULL)) {
		return 0;
	}
	int l = strlen(a[0]);
	if (l != strlen(b[0])) return 0;

	for (int i = 0; i < n; i++) {
		for (int j = 0; j < l; j++) {
			if (a[i][j] != b[i][j]) return 0;
		}
	}
	return 1;
}

long count_occ(char** s, int n)
{
	int l = strlen(s[0]);
	long c = 0;
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < l; j++) {
			c += (s[i][j] == OCC);
		}
	}
	return c;
}

char** step_state(char** s, int n)
{
	char** ns = calloc(1,sizeof(char*)*n);
	int l = strlen(s[0]);
	int adj = 0, c, d;

	for (int i = 0; i < n; i++) {
		ns[i] = calloc(1,sizeof(char)*l+1);
		for (int j = 0; j < l; j++) {
			if (s[i][j] == FLOOR) {
				ns[i][j] = s[i][j]; continue;
			}
			adj = 0;

			// NW
			c = i; d = j;
			while (c > 0 && d > 0) if (s[--c][--d] != FLOOR) break;
			adj += (i != c || j != d) && s[c][d] == OCC;

			// NE
			c = i; d = j;
			while (c > 0 && d < l-1) if (s[--c][++d] != FLOOR) break;
			adj += (i != c || j != d) && s[c][d] == OCC;

			// SW
			c = i; d = j;
			while (c < n-1 && d > 0) if (s[++c][--d] != FLOOR) break;
			adj += (i != c || j != d) && s[c][d] == OCC;

			// SE
			c = i; d = j;
			while (c < n-1 && d < l-1) if (s[++c][++d] != FLOOR) break;
			adj += (i != c || j != d) && s[c][d] == OCC;

			// N
			c = i; d = j;
			while (c > 0) if (s[--c][d] != FLOOR) break;
			adj += (i != c || j != d) && s[c][d] == OCC;

			// S
			c = i; d = j;
			while (c < n-1) if (s[++c][d] != FLOOR) break;
			adj += (i != c || j != d) && s[c][d] == OCC;

			// W
			c = i; d = j;
			while (d > 0) if (s[c][--d] != FLOOR) break;
			adj += (i != c || j != d) && s[c][d] == OCC;

			// E
			c = i; d = j;
			while (d < l-1) if (s[c][++d] != FLOOR) break;
			adj += (i != c || j != d) && s[c][d] == OCC;

			//printf("%c %3d %3d %3d\n", s[i][j], i, j, adj);

			if (s[i][j] == EMPTY && adj == 0) {
				ns[i][j] = OCC;
			} else if (s[i][j] == OCC && adj >= 5) {
				ns[i][j] = EMPTY;
			} else {
				ns[i][j] = s[i][j];
			}
		}
	}

	return ns;
}

long count_steps(char** s, int n)
{
	char** prev = NULL;
	char** curr = s;
	long n_states = 0;
	while (!states_equal(prev, curr, n)) {
		//print_state(curr, n);
		if (prev) ffree((void**)prev, n);
		prev = curr;
		curr = step_state(prev, n);
		n_states++;
	}

	DUMP("%ld", n_states);

	long res = count_occ(curr, n);
	ffree((void**)prev, n);
	ffree((void**)curr, n);
	return res;
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** init_state = load_file(FN, 1 << 14, &n);
	assert(init_state);

	long res = count_steps(init_state, n);
	DUMP("%ld", res);
#if TEST
	assert(res == 26);
#endif

	return 0;
}
