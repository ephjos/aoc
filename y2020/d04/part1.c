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

	// Read data
	int
		byr = 0,
		iyr = 0,
		eyr = 0,
		hgt = 0,
		hcl = 0,
		ecl = 0,
		pid = 0,
		cid = 0;
	char buf[1024] = {0};
	char temp[512] = {0};
	char* p;
	int valid = 0;
	int n = 0;

	while(fgets(buf, 512, fp)) {
		if (strcmp(buf, "\n") == 0) {
			valid += byr && iyr && eyr && hgt && hcl && ecl && pid;
			byr = 0; iyr = 0; eyr = 0; hgt = 0;
			hcl = 0; ecl = 0; pid = 0; cid = 0;
		}

		p = strtok(buf, ":");
		while (p != NULL) {
			if (strcmp(p, "byr") == 0) { byr ++; }
			if (strcmp(p, "iyr") == 0) { iyr ++; }
			if (strcmp(p, "eyr") == 0) { eyr ++; }
			if (strcmp(p, "hgt") == 0) { hgt ++; }
			if (strcmp(p, "hcl") == 0) { hcl ++; }
			if (strcmp(p, "ecl") == 0) { ecl ++; }
			if (strcmp(p, "pid") == 0) { pid ++; }
			if (strcmp(p, "cid") == 0) { cid ++; }
			p = strtok(NULL, " ");
			p = strtok(NULL, ":");
		}
	}

	// Need this to get correct answer...
	valid += byr && iyr && eyr && hgt && hcl && ecl && pid;
	byr = 0; iyr = 0; eyr = 0; hgt = 0;
	hcl = 0; ecl = 0; pid = 0; cid = 0;

	printf("valid = %d\n", valid);

	// Call function on data here

	// Cleanup
	fclose(fp);

	return 0;
}
