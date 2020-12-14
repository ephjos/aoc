#include <math.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

typedef struct vec2_t vec2;
typedef struct vec2_t {
	long x, y;
} vec2;

typedef struct boat_t boat;
typedef struct boat_t {
	vec2* pos;
	vec2* dir;
} boat;

void free_boat(boat* b)
{
	free(b->pos); free(b->dir);
	free(b);
	return;
}

vec2* init_vec2()
{
	vec2* v = malloc(sizeof(vec2));
	v->x = 0L; v->y = 0L;
	return v;
}

boat* init_boat()
{
	boat* b = malloc(sizeof(boat));
	b->pos = init_vec2();
	b->dir = init_vec2();
	b->dir->x = 1L;
	return b;
}

double deg2rad(int deg)
{
	return deg * (M_PI/180.0);
}

void rotate_vec2(vec2* v, int deg)
{
	double rad = deg2rad(deg);
	long nx = (cos(rad)*v->x) - (sin(rad)*v->y);
	long ny = (sin(rad)*v->x) + (cos(rad)*v->y);
	v->x = nx; v->y = ny;
	return;
}

void add_vec2(vec2* dest, vec2* src)
{
	dest->x += src->x; dest->y += src->y;
	return;
}

vec2* scale_vec2(vec2* v, int s)
{
	vec2* r = init_vec2();
	r->x = v->x*s; r->y = v->y*s;
	return r;
}

long l1_vec2(vec2* v)
{
	return ABS(v->x) + ABS(v->y);
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines;

	if (argc == 3 &&
			(strcmp(argv[1], "-f") == 0 ||
			 strcmp(argv[1], "--file") == 0)) {
		lines = load_file(argv[2], 0, &n);
	} else {
		lines = load_file("./input", 0, &n);
	}

	char c;
	int v;
	boat* b = init_boat();
	vec2* move;
	for (int i = 0; i < n; i++) {
		sscanf(lines[i], "%c%d\n", &c, &v);
		switch(c) {
			case 'N':
				b->pos->y += v;
				break;
			case 'E':
				b->pos->x += v;
				break;
			case 'S':
				b->pos->y -= v;
				break;
			case 'W':
				b->pos->x -= v;
				break;
			case 'F':
				move = scale_vec2(b->dir, v);
				add_vec2(b->pos, move);
				free(move);
				break;
			case 'L':
				rotate_vec2(b->dir, v);
				break;
			case 'R':
				rotate_vec2(b->dir, -1 * v);
				break;
			default:
				break;
		}
		printf("%c %6d (%6ld, %6ld) (%6ld, %6ld)\n",
				c, v, b->dir->x, b->dir->y, b->pos->x, b->pos->y);
	}

	printf("dist = %8ld\n", l1_vec2(b->pos));

	free_boat(b);
	ffree((void**)lines, n);
	return 0;
}
