#define YEAR 2015
#define DAY 1

#include "../shared_2015.h"

// In ASCII ( is 40, ) is 41. This maps c-40 to the corresponding floor move
const int CHAR_TO_MOVE[2] = {1, -1};

int main(const int argc, const char *argv[]) {
	struct input_file file = get_input_file();

	int64_t floor = 0;
	uint64_t crossed_to_basement = 0;
	for (uint64_t i = 0; i < file.len; i++) {
		floor += CHAR_TO_MOVE[file.bytes[i]-40];
		if (floor < 0 && !crossed_to_basement) {
			crossed_to_basement = i+1;
		}
	}

	part_1("%ld", floor);
	part_2("%lu", crossed_to_basement);

	free_input_file(&file);
	return 0;
}
