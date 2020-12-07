#include <stdlib.h>
#include <stdio.h>
#include <string.h>

void ehandler(char* s, int exit_code)
{
	fprintf(stderr, "%s\n", s);
	exit(exit_code);
}

long num_lines(FILE* fp)
{
	long count = 0;
	char buf[128] = {0};
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

	// Read data
	long* vs = (long*)calloc(1, n*sizeof(long));
	char buf[20] = {0};
	for (int i = 0; fscanf(fp,"%s\n", buf) == 1; i++) {
		vs[i] = atol(buf);
  }

	// Call function on data here

	// Cleanup
	free(vs);
	fclose(fp);

	return 0;
}
