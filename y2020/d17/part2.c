#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

const int DEAD = 0;
const int ALIVE = 1;

const char DEADC = '.';
const char ALIVEC = '#';

int**** new_space(int N)
{
	int**** space = malloc(sizeof(int***)*N);
	for (int i = 0; i < N; i++) {
		space[i] = malloc(sizeof(int**)*N);
		for (int j = 0; j < N; j++) {
			space[i][j] = malloc(sizeof(int*)*N);
			for (int k = 0; k < N; k++) {
				space[i][j][k] = malloc(sizeof(int)*N);
				for (int l = 0; l < N; l++) {
					space[i][j][k][l] = DEAD;
				}
			}
		}
	}
	return space;
}

void free_space(int**** space, int N)
{
	for (int i = 0; i < N; i++) {
		for (int j = 0; j < N; j++) {
			for (int k = 0; k < N; k++) {
				free(space[i][j][k]);
			}
			free(space[i][j]);
		}
		free(space[i]);
	}
	free(space);
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

	int o = 24; // padding around input slice
	int N = n+o; // add padding
	int c = N / 2; // center

	// Initiaize space
	int**** space = new_space(N);

	// Initiaize slice
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			// Center in center slice
			space[c][c][i+(o/2)][j+(o/2)] = lines[i][j] == ALIVEC;
		}
	}

	// Print slice
	/*
	for (int i = 0; i < N; i++) {
		for (int j = 0; j < N; j++) {
			printf("%d ", space[c][i][j]);
		}
		if (i == N-1) printf("\n");
	}
	*/


	// Run steps
	int alive_neighbors;
	int**** n_space, ****p_space = space;
	for (int s = 0; s < 6; s++) {
		n_space = new_space(N);
		for (int i = 0; i < N; i++) {
			for (int j = 0; j < N; j++) {
				for (int k = 0; k < N; k++) {
					for (int l = 0; l < N; l++) {
						alive_neighbors = 0;
						for (int x = -1; x <= 1; x++) {
							for (int y = -1; y <= 1; y++) {
								for (int z = -1; z <= 1; z++) {
									for (int w = -1; w <= 1; w++) {
									if (x == y && y == z && z == w && w == 0) continue; // skip self
									if (IS_BETWEEN(i+x, 0, N-1) && IS_BETWEEN(j+y, 0, N-1) &&
											IS_BETWEEN(k+z, 0, N-1) && IS_BETWEEN(l+w, 0, N-1)) {
										alive_neighbors += p_space[i+x][j+y][k+z][l+w];
									}
								}
							}
							}
						}
						n_space[i][j][k][l] = ((p_space[i][j][k][l] && (alive_neighbors == 2 || alive_neighbors == 3))
									|| (!p_space[i][j][k][l] && alive_neighbors == 3));
					}
				}
			}
		}

		// Print slice
		/*
		for (int i = 0; i < N; i++) {
			for (int j = 0; j < N; j++) {
				printf("%d ", n_space[c][i][j]);
			}
			printf("\n");
			if (i == N-1) printf("\n");
		}
		*/
		free_space(p_space, N);
		p_space = n_space;
	}

	// Count alive
	int sum = 0;
	for (int i = 0; i < N; i++) {
		for (int j = 0; j < N; j++) {
			for (int k = 0; k < N; k++) {
				for (int l = 0; l < N; l++) {
					sum += n_space[i][j][k][l];
				}
			}
		}
	}

	DUMP("%d", sum);

	// Cleanup
	free_space(n_space, N);
	FFREE(lines, n);
	return 0;
}
