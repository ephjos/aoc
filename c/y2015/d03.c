#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#define YEAR 2015
#define DAY 3

#include "../shared_2015.h"

int main(const int argc, const char *argv[]) {
	struct input_file file = get_input_file();

	// Part 1
	uint32_t seen_len = 0;
	struct point seen_points[file.len];
	uint32_t seen_counts[file.len];

	uint64_t houses = 0;

	struct point santa = { 0, 0 };
	seen_points[seen_len] = santa;
	seen_counts[seen_len] = 1;
	seen_len++;

	// Part 2
	uint32_t seen_len_2 = 0;
	struct point seen_points_2[file.len];
	uint32_t seen_counts_2[file.len];

	uint64_t houses_2 = 0;

	struct point santa_2 = { 0, 0 };
	struct point robo_2 = { 0, 0 };
	struct point *curr_2;
	seen_points_2[seen_len_2] = santa_2;
	seen_counts_2[seen_len_2] = 1;
	seen_len_2++;

	for (uint64_t i = 0; i < file.len; i++) {
		if (i % 2 == 0) {
			curr_2 = &santa_2;
		} else {
			curr_2 = &robo_2;
		}

		switch (file.bytes[i]) {
			case '>':
				santa.x += 1;
				curr_2->x += 1;
				break;
			case '<':
				santa.x -= 1;
				curr_2->x -= 1;
				break;
			case '^':
				santa.y += 1;
				curr_2->y += 1;
				break;
			case 'v':
				santa.y -= 1;
				curr_2->y -= 1;
				break;
			default:
				break;
		}

		// Part 1
		uint32_t j;
		for (j = seen_len-1; j > 0; j--) {
			if (POINT_EQUALS(seen_points[j], santa)) {
				break;
			}
		}

		if (j == 0) {
			seen_points[seen_len] = santa;
			seen_counts[seen_len] = 1;
			seen_len++;
			houses++;
		} else {
			seen_counts[j]++;
		}

		// Part 2
		for (j = seen_len_2-1; j > 0; j--) {
			if (POINT_EQUALS(seen_points_2[j], (*curr_2))) {
				break;
			}
		}

		if (j == 0) {
			seen_points_2[seen_len_2] = (*curr_2);
			seen_counts_2[seen_len_2] = 1;
			seen_len_2++;
			houses_2++;
		} else {
			seen_counts_2[j]++;
		}
	}

	part_1("%"PRIu64, houses);
	part_2("%"PRIu64, houses_2);

	free_input_file(&file);
	return 0;
}
