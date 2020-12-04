#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "../include/cfx.h"

void ehandler(char* s, int exit_code)
{
	fprintf(stderr, "%s\n", s);
	exit(exit_code);
}

long mult_sum_2020(long* vs, long n)
{
	for (long i = 0; i < n; i++) {
		for (long j = 0; j < n; j++) {
			if (vs[i] + vs[j] == 2020) {
				return vs[i]*vs[j];
			}
		}
	}
	return 0l;
}

long num_lines(FILE* fp)
{
	long count = 0;
	char buf[20] = {0};
	fseek(fp, 0, 0);
	for (; fscanf(fp,"%s\n", buf) == 1; ) {
		count++;
  }

	fseek(fp, 0, 0);
	return count;
}

int main(int argc, char *argv[])
{
	FILE* fp = fopen("./input", "r");
	if (fp == NULL) {
		ehandler("Could not load ./input file", 1);
	}

	long n = num_lines(fp);
	long* vs = (long*)calloc(1, n*sizeof(long));

	char buf[20] = {0};
	for (int i = 0; fscanf(fp,"%s\n", buf) == 1; i++) {
		vs[i] = atol(buf);
  }

	long res = mult_sum_2020(vs, n);
	printf("%ld\n", res);

	free(vs);
	fclose(fp);

	return 0;
}
