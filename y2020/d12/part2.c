#include <math.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

typedef struct vec2_t vec2;
typedef struct vec2_t {
	double x, y;
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
	double nx = (cos(rad)*v->x) - (sin(rad)*v->y);
	double ny = (sin(rad)*v->x) + (cos(rad)*v->y);
	v->x = nx; v->y = ny;
	return;
}

void addi_vec2(vec2* dest, vec2* src)
{
	dest->x += src->x; dest->y += src->y;
	return;
}

vec2* scale_vec2(vec2* v,  double s)
{
	vec2* r = init_vec2();
	r->x = v->x*s; r->y = v->y*s;
	return r;
}

vec2* sub_vec2(vec2* a, vec2* b)
{
	vec2* r = init_vec2();
	r->x = a->x - b->x; r->y = a->y - b->y;
	return r;
}

double l1_vec2(vec2* v)
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
	vec2* b = init_vec2();
	vec2* wp = init_vec2();
	wp->x = 10; wp->y = 1;
	vec2 *dist, *move;
	for (int i = 0; i < n; i++) {
		sscanf(lines[i], "%c%d\n", &c, &v);
		switch(c) {
			case 'N':
				wp->y += v;
				break;
			case 'E':
				wp->x += v;
				break;
			case 'S':
				wp->y -= v;
				break;
			case 'W':
				wp->x -= v;
				break;
			case 'F':
				move = scale_vec2(wp, v);
				addi_vec2(b, move);
				free(move);
				break;
			case 'L':
				rotate_vec2(wp, v);
				break;
			case 'R':
				rotate_vec2(wp, -1 * v);
				break;
			default:
				break;
		}
		printf("%c %6d (%2f, %2f) (%2f, %2f)\n",
				c, v, wp->x, wp->y, b->x, b->y);
	}

	printf("dist = %2f\n", l1_vec2(b));

	free(wp);
	free(b);
	ffree((void**)lines, n);
	return 0;
}
