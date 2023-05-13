//
// This header file is used by defining the YEAR and DAY macros as the integers
// that correspond to the problem date.
//

#ifndef __SHARED_2015__
#define __SHARED_2015__

#ifndef YEAR
#error YEAR is not defined // Forces compilation to fail if YEAR is not defined
#define YEAR 0 // Makes the rest of this code pass checks in isolation, never
							 // hit since this branch exits with the above exit
#endif

#ifndef DAY
#error DAY is not defined // Forces compilation to fail if DAY is not defined
#define DAY 0 // Makes the rest of this code pass checks in isolation, never
							 // hit since this branch exits with the above exit
#endif

// =============================================================================
// Includes
// =============================================================================
#include <assert.h>
#include <ctype.h>
#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>


// =============================================================================
// Number XMacros
// =============================================================================

// Lists all of the single-word name number types
#define NUMBER_TYPES(F) \
	F(char, _) \
	F(short, _) \
	F(int, _) \
	F(long, _) \
	F(float, _) \
	F(double, _) \
	F(uint8_t, _) \
	F(uint16_t, _) \
	F(uint32_t, _) \
	F(uint64_t, _) \
	F(int8_t, _) \
	F(int16_t, _) \
	F(int32_t, _) \
	F(int64_t, _) \

// Define a qsort compatible comparator for each number type
#define NUMBER_COMPARATOR(n, ...) \
int compare_##n(const void *a, const void *b) { \
	n cast_a = *((n*)a); \
	n cast_b = *((n*)b); \
	return (cast_a > cast_b) - (cast_a < cast_b); \
} \

NUMBER_TYPES(NUMBER_COMPARATOR)

// Define a simple qsort wrapper for each number type
#define NUMBER_QSORT(n, ...) \
void qsort_##n(n *arr, uint32_t nmemb) { \
	qsort(arr, nmemb, sizeof(n), compare_##n); \
} \

NUMBER_TYPES(NUMBER_QSORT)

// =============================================================================
// Macros
// =============================================================================
#define eprintf( format, ... ) fprintf(stderr, format, __VA_ARGS__)
#define part_1( format, ... ) fprintf(stdout, "%d.%02d.1: " format "\n", YEAR, DAY, __VA_ARGS__)
#define part_2( format, ... ) fprintf(stdout, "%d.%02d.2: " format "\n", YEAR, DAY, __VA_ARGS__)

#define DDUMP( fmt, ... ) \
	fprintf(stderr, "%s:%d:%s(): " fmt "\n", __FILE__, \
			__LINE__, __func__, ##__VA_ARGS__); \

#define DUMP( fmt, val ) \
	printf("%s = " fmt "\n", #val, val);

#undef MAX
#define MAX( a, b ) ((a) > (b) ? (a) : (b))

#undef MIN
#define MIN( a, b ) ((a) > (b) ? (b) : (a))

#undef ABS
#define ABS( x ) (((x) < 0) ? (-x) : (x))

#define IS_BETWEEN( x, a, b ) \
	((unsigned char)((x) >= (a) && (x) <= (b)))

#define XIS_BETWEEN( x, a, b ) \
	((unsigned char)((x) > (a) && (x) < (b)))


// =============================================================================
// Input File
// =============================================================================
#define INPUT_FILENAME_BUF_SIZE 24
struct input_file {
	uint64_t len;
	char filename[INPUT_FILENAME_BUF_SIZE];
	char *bytes;
};


void free_input_file(struct input_file *i);
inline void free_input_file(struct input_file *i) {
	free(i->bytes);
}

// Load an input file for a given day
struct input_file get_input_file();
inline struct input_file get_input_file() {
	struct input_file i;

	// Build the file path
	int32_t n = snprintf(i.filename, INPUT_FILENAME_BUF_SIZE, "../../inputs/y%d/d%02d", YEAR, DAY);
	if (n < 0) {
		eprintf("Could not write filename with year %d and day %d\n", YEAR, DAY);
		exit(1);
	} else if (n > INPUT_FILENAME_BUF_SIZE) {
		eprintf("Filename buffer was too small, had %d but needed %d", INPUT_FILENAME_BUF_SIZE, n);
		exit(1);
	}

	// Open the file for reading
	FILE *fp = fopen(i.filename, "r");
	if (fp == NULL) {
		eprintf("Could not open file \"%s\" for reading\n", i.filename);
		exit(1);
	}

	// Get the total size of the file (in bytes)
	fseek(fp, 0, SEEK_END);
	i.len = ftell(fp);
	fseek(fp, 0, SEEK_SET);

	// Load file into memory
	i.bytes = (char*)malloc(i.len+1);
	fread(i.bytes, i.len, 1, fp);
	i.bytes[i.len] = '\0';

	// Cleanup
	fclose(fp);

	return i;
}

// =============================================================================
// Points
// =============================================================================
struct point {
	int32_t x;
	int32_t y;
};

#define POINT_EQUALS(a, b) (a.x == b.x && a.y == b.y)

#endif
