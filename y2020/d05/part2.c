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

void parse_passes(char** vs, long n)
{
	int* x = (int*)malloc(sizeof(int)*n);
	int a,b,c,d,max=0,id=0,row=0,col=0;
	for (int i = 0; i<n; i++) {
		a=0;b=127;c=0;d=7;
		for (int j = 0; j<7; j++) {
			if (vs[i][j] == 'F') {
				b -= (b-a)/2 + 1;
			} else {
				a += (b-a)/2 + 1;
			}
		}
		row = a;
		for (int j = 7; j<10; j++) {
			if (vs[i][j] == 'L') {
				d -= (d-c)/2 + 1;
			} else {
				c += (d-c)/2 + 1;
			}
		}
		col = c;

		id = row*8 + col;
		if (id > max) max = id;
		x[i] = id;
	}

	int* s = (int*)calloc(1, sizeof(int)*max+1);
	int q;
	for (int i = 0; i<n; i++) {
		q = x[i];
		s[q] = 1;
	}

	for (int i = 0; i<n; i++) {
		q = x[i];
		if (q >= 2 && s[q-1] == 0 && s[q-2] == 1) {
			printf("my seat = %d\n", q-1);
		} else if (q < n-2 && s[q+1] == 0 && s[q+2] == 1) {
			printf("my seat = %d\n", q+1);
		}
	}
}

int main(int argc, char *argv[])
{
	FILE* fp = fopen("./input", "r");
	if (fp == NULL) {
		ehandler("Could not load ./input file", 1);
	}

	long n = num_lines(fp);

	// Read data
	char** vs = (char**)calloc(1, n*sizeof(char*));
	char buf[256] = {0};
	for (int i = 0; fscanf(fp,"%s\n", buf) == 1; i++) {
		vs[i] = (char*)calloc(1, 512*sizeof(char));
		strcpy(vs[i], buf);
  }

	// Call function on data here
	parse_passes(vs, n);

	// Cleanup
	free(vs);
	fclose(fp);

	return 0;
}
