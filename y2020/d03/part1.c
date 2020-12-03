#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define DEBUG 1
#define DDUMP(fmt, ...) \
        do { if (DEBUG) fprintf(stderr, "%s:%d:%s(): " fmt "\n", __FILE__, \
                                __LINE__, __func__, __VA_ARGS__); } while (0)
#define DUMP(fmt, val) \
	printf("%s = " fmt "\n", #val, val);

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

int ctos(char** cs, int height, int width, int right, int down)
{
	int row = 0;
	int col = 0;
	int count = 0;
	for (; row<height; row+=down, col += right) {
		count += cs[row][col%width] == '#';
	}

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
	int height = n;
	int width = 0;
	char** cs = (char**)calloc(1, n*sizeof(char*));
	#define BUF_SIZE 512
	char buf[BUF_SIZE] = {0};
	for (int i = 0; fscanf(fp,"%s\n", buf) == 1; i++) {
		cs[i] = (char*)calloc(1, sizeof(char)*BUF_SIZE);
		strncpy(cs[i], buf, BUF_SIZE);
		width = strlen(buf);
  }

	// Call function on data here
	int count = ctos(cs, height, width, 3, 1);
	DUMP("%d", count);

	// Cleanup
	free(cs);
	fclose(fp);

	return 0;
}
