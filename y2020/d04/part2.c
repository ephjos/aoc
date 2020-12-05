#include <regex.h>
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
	char* v;
	int valid = 0;
	int n = 0;
	regex_t hcl_re;
  if (regcomp(
				&hcl_re,
				".*#[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f].*",
				0) != 0) {
		exit(1);
	}

	regex_t ecl_re;
  if (regcomp(
				&ecl_re,
				".*(amb|blu|brn|gry|grn|hzl|oth).*",
				REG_EXTENDED) != 0) {
		exit(1);
	}

	regex_t pid_re;
  if (regcomp(
				&pid_re,
				".*[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9].*",
				0) != 0) {
		exit(1);
	}

	while(fgets(buf, 512, fp)) {
		if (strcmp(buf, "\n") == 0) {
			printf("%d %d %d %d %d %d %d\n",
					byr, iyr, eyr, hgt, hcl, ecl, pid);
			valid += byr && iyr && eyr && hgt && hcl && ecl && pid;
			byr = 0; iyr = 0; eyr = 0; hgt = 0;
			hcl = 0; ecl = 0; pid = 0; cid = 0;
			continue;
		}

		p = strtok(buf, ":");
		while (p != NULL) {
			v = strtok(NULL, " ");
			int m = strlen(v);
			if (v[m-1] == '\n') { v[m-1] = '\0'; }
			printf("p = '%s' | v = '%s'\n", p, v);
			if (strcmp(p, "byr") == 0) {
				int x = atoi(v);
				if (1920 <= x && x <= 2002) {
					byr++;
				}
			}
			if (strcmp(p, "iyr") == 0) {
				int x = atoi(v);
				if (2010 <= x && x <= 2020) {
					iyr++;
				}
			}
			if (strcmp(p, "eyr") == 0) {
				int x = atoi(v);
				if (2020 <= x && x <= 2030) {
					eyr++;
				}
			}
			if (strcmp(p, "hgt") == 0) {
				int n = strlen(v);
				if (v[n-2] == 'c' && v[n-1] == 'm') {
					v[n-2] = '\0'; v[n-1] = '\0';
					int x = atoi(v);
					printf("x = %d\n", x);
					if (150 <= x && x <= 193) {
						hgt++;
					}
					v[n-2] = 'c'; v[n-1] = 'm';
				}
				if (v[n-2] == 'i' && v[n-1] == 'n') {
					v[n-2] = '\0'; v[n-1] = '\0';
					int x = atoi(v);
					printf("  x = %d\n", x);
					if (59 <= x && x <= 76) {
						hgt++;
					}
					v[n-2] = 'i'; v[n-1] = 'n';
				}
				printf("  hgt = %d\n", hgt);
			}
			if (strcmp(p, "hcl") == 0) {
				if (regexec(&hcl_re, v, 0, NULL, 0) == 0) {
					hcl++;
				}
			}
			if (strcmp(p, "ecl") == 0) {
				if (regexec(&ecl_re, v, 0, NULL, 0) == 0) {
					ecl++;
				}
			}
			if (strcmp(p, "pid") == 0) {
				if (regexec(&pid_re, v, 0, NULL, 0) == 0) {
					pid++;
				}
			}
			if (strcmp(p, "cid") == 0) { cid ++; }
			v[m-1] = '\n';
			p = strtok(NULL, ":");
		}
	}

	// Had to remove to get correct answer...

	printf("valid = %d\n", valid);

	// Call function on data here

	// Cleanup
	fclose(fp);

	return 0;
}
