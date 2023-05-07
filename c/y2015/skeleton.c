#define YEAR 2015
#define DAY 2

#include "../shared_2015.h"

int main(const int argc, const char *argv[]) {
	struct input_file file = get_input_file();

	printf("%s\n", file.bytes);

	free_input_file(&file);
	return 0;
}
