#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "../include/cge.h"

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines = load_file("./input", 0, &n);



	ffree((void**)lines, n);
	return 0;
}
