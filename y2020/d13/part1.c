#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

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

	long el;
	if (sscanf(lines[0], "%ld\n", &el) != 1) ERR("Couldnt scan\n");

	int m;
	char** toks = split(lines[1], ",", &m);

	long min = 1<<12, min_id = 1<<12;
	long id;
	long temp;
	short new_min;
	for (int i = 0; i < m; i++) {
		if (toks[i][0] == 'x') continue;
		sscanf(toks[i], "%ld\n", &id);
		printf("el = %6ld id = %6ld mod = %6ld\n", el, id,  el % id);
		temp = id - (el % id);
		new_min = temp < min;
		min = (min * !new_min) + (temp * new_min);
		min_id = (min_id * !new_min) + (id * new_min);
	}

	printf("wait = %6ld id = %6ld res = %6ld\n", min, min_id, min*min_id);

	FFREE(toks, m);
	FFREE(lines, n);
	return 0;
}
