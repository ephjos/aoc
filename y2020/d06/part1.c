#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

void ehandler(char* s, int exit_code)
{
	fprintf(stderr, "%s\n", s);
	exit(exit_code);
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines = load_file("./input", 0, &n);
	int* cs = (int*)calloc(1,sizeof(int)*26);
	int m = 0;
	int sum = 0;
	int tsum = 0;

	// Call function on data here
	for (int i = 0; i < n; i++) {
		if (strcmp(lines[i], "") == 0) {
			sum = 0;
			for (int j = 0; j < 26; j++) {
				sum += cs[j];
			}
			tsum += sum;
			m = MAX(m, sum);
			free(cs);
			cs = (int*)calloc(1,sizeof(int)*26);
			continue;
		}
		for (int j = 0; j < strlen(lines[i]); j++) {
			cs[lines[i][j]-97] = 1;
		}
	}
	sum = 0;
	for (int j = 0; j < 26; j++) {
		sum += cs[j];
	}
	tsum += sum;
	m = MAX(m, sum);
	free(cs);

	DUMP("%d", tsum);

	// Cleanup
	ffree((void**)lines, n);

	return 0;
}
