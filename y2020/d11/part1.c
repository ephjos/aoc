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
	char* ts[5][10] = {
		{
			"#.##.##.##",
			"#######.##",
			"#.#.#..#..",
			"####.##.##",
			"#.##.##.##",
			"#.#####.##",
			"..#.#.....",
			"##########",
			"#.######.#",
			"#.#####.##",
		},
		{
			"#.LL.L#.##",
			"#LLLLLL.L#",
			"L.L.L..L..",
			"#LLL.LL.L#",
			"#.LL.LL.LL",
			"#.LLLL#.##",
			"..L.L.....",
			"#LLLLLLLL#",
			"#.LLLLLL.L",
			"#.#LLLL.##",
		},
		{
			"#.##.L#.##",
			"#L###LL.L#",
			"L.#.#..#..",
			"#L##.##.L#",
			"#.##.LL.LL",
			"#.###L#.##",
			"..#.#.....",
			"#L######L#",
			"#.LL###L.L",
			"#.#L###.##",
		},
		{
			"#.#L.L#.##",
			"#LLL#LL.L#",
			"L.L.L..#..",
			"#LLL.##.L#",
			"#.LL.LL.LL",
			"#.LL#L#.##",
			"..L.L.....",
			"#L#LLLL#L#",
			"#.LLLLLL.L",
			"#.#L#L#.##",
		},
		{
			"#.#L.L#.##",
			"#LLL#LL.L#",
			"L.#.L..#..",
			"#L##.##.L#",
			"#.#L.LL.LL",
			"#.#L#L#.##",
			"..L.L.....",
			"#L#L##L#L#",
			"#.LLLLLL.L",
			"#.#L#L#.##",
		}};
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
	int adj = 0;

	for (int i = 0; i < n; i++) {
		ns[i] = calloc(1,sizeof(char)*l+1);
		for (int j = 0; j < l; j++) {
			if (s[i][j] == FLOOR) {
				ns[i][j] = s[i][j]; continue;
			}
			adj = 0;
			adj = ((i > 0 && s[i-1][j] == OCC) +
						 (i > 0 && j > 0 && s[i-1][j-1] == OCC) +
						 (i > 0 && j < l-1 && s[i-1][j+1] == OCC) +
						 (i < n-1 && s[i+1][j] == OCC) +
						 (i < n-1 && j > 0 && s[i+1][j-1] == OCC) +
						 (i < n-1 && j < l-1 && s[i+1][j+1] == OCC) +
						 (j > 0 && s[i][j-1] == OCC) +
						 (j < l-1 && s[i][j+1] == OCC));

			if (s[i][j] == EMPTY && adj == 0) {
				ns[i][j] = OCC;
			} else if (s[i][j] == OCC && adj >= 4) {
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
#if TEST
		if (n_states < 5) {
			int eq = states_equal(curr, ts[n_states], n);
			if (!eq) {
				DUMP("%ld", n_states);
				print_state(curr, n);
				print_state(ts[n_states], n);
				assert(eq);
			}
		}
#endif
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
#if TEST
	assert(res == 37);
#endif
	DUMP("%ld", res);

	return 0;
}
