#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

long long modinv(long long a, long long b)
{
	long long
		old_r = a, r = b,
		old_s = 1, s = 0,
		old_t = 0, t = 1;

	long long q, tr, ts, tt;
	while (r != 0) {
		q = old_r / r;
		tr = old_r; old_r = r; r = tr - q * r;
		ts = old_s; old_s = s; s = ts - q * s;
		tt = old_t; old_t = t; t = tt - q * t;
	}

	if (old_s < 0) return b + old_s;
	return old_s;
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

	long long el;
	if (sscanf(lines[0], "%lld\n", &el) != 1) ERR("Coulldnt scan\n");

	int m, k = 0;
	char** toks = split(lines[1], ",", &m);
	long long nn, N = 1;
	long long* as = malloc(sizeof(long long)*m);
	long long* ns = malloc(sizeof(long long)*m);
	for (int i = 0; i < m; i++) {
		if (toks[i][0] == 'x') continue;
		printf("%d %s\n", i, toks[i]);
		sscanf(toks[i], "%lld\n", &nn);
		N *= nn;
		as[k] = nn - i; ns[k] = nn;
		k++;
	}
	as = realloc(as, sizeof(long long)*k); ns = realloc(ns, sizeof(long long)*k);

	long long x = 0,y,z;
	for (int i = 0; i < k; i++) {
		y = N / ns[i];
		z = modinv(y, ns[i]);
		x += as[i]*y*z;
	}

	printf("%6lld %6lld\n", x, x % N);

	free(as);
	free(ns);
	FFREE(toks, m);
	FFREE(lines, n);
	return 0;
}
