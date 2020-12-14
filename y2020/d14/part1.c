#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"
#include "clist.h"

#define MEM_ISIZE 8192

void print_bin(ulong x)
{
	//ulong l = sizeof(x)*8;
	ulong l = 36;
	for (long i = l-1; i >= 0; i--) {
		printf("%lu", (x >> i) & 1);
	}
	printf("\n");
}

void parse_mask(char* line, ulong* m1, ulong* m2)
{
	char buf[48] = {0};
	sscanf(line, "mask = %s\n", buf);
	#define MASK_LEN 36
	ulong t1 = 0;
	ulong t2 = 0;

	for(int i = 0; i < MASK_LEN; i++) {
		t1 = (t1 | (buf[i] == 'X')) << (i != MASK_LEN-1);
		t2 = (t2 | (buf[i] == '1')) << (i != MASK_LEN-1);
	}

	*m1 = t1;
	*m2 = t2;
}

int main(int argc, char *argv[])
{
	int ln = 0;
	char** lines;

	if (argc == 3 &&
			(strcmp(argv[1], "-f") == 0 ||
			 strcmp(argv[1], "--file") == 0)) {
		lines = load_file(argv[2], 0, &ln);
	} else {
		lines = load_file("./input", 0, &ln);
	}

	ulong_list* mem = new_ulong_list(MEM_ISIZE);
	ulong m1, m2 = 0;
	ulong addr, val;

	for (int i = 0; i < ln; i++) {
		if (lines[i][1] == 'a') { // mask
			parse_mask(lines[i], &m1, &m2);
		} else { // mem
			sscanf(lines[i], "mem[%ld] = %ld\n", &addr, &val);
			set_ulong_list(mem, addr, (val & m1) | m2);
		}
	}

	ulong sum = 0;
	for (int i = 0; i < mem->size; i++) {
		sum += get_ulong_list(mem, i);
	}

	DUMP("%lu", sum);

	free_ulong_list(mem);
	FFREE(lines, ln);
	return 0;
}
