#include <stdio.h>

#include "../include/cge.h"
#include "../include/cfx.h"

void draw_seats(int px, int py, int** ss)
{
	int w = 10, h = 10;
	for (int i = 0; i < 8; i++) {
		if (i == 4) py += h;
		for (int j = 0; j < 128; j++) {
			cfx_color(
					(30*!ss[i][j]) + (96*ss[i][j]),
					(30*!ss[i][j]) + (241*ss[i][j]),
					(30*!ss[i][j]) + (0*ss[i][j]));
			cfx_draw_rectangle(px+(j*w), py+(i*w), w, h, ss[i][j]);
		}
	}
}

void storc(char* s, int* r, int* c)
{
	int tr = 0, tc = 0;
	for (int i = 0; i < 7; i++) {
		tr += s[i] == 'B';
		tr <<= 1;
	}
	tr >>= 1;
	for (int i = 7; i < 10; i++) {
		tc += s[i] == 'R';
		tc <<= 1;
	}
	tc >>= 1;
	*r = tr;
	*c = tc;
}

void get_mine(int* ids, int* fc, int* fr)
{
	for (int i = 1; i < (8*128)-1; i++) {
		if (ids[i] == 0 && ids[i-1] == 1 && ids[i+1] == 1) {
			*fc = i & 0b0000000111;
			*fr = i >> 3;
		}
	}
}

int main()
{
	int xsize = 1300;
	int ysize = 150;

	char c;
	int x,y;

	// Open a new window for drawing.
	//cfx_open(xsize, ysize, 60, "vis.mp4");
	cfx_open(xsize, ysize, 60, NULL);

	int i = 0;
	int** ss = (int**)calloc(1, sizeof(int*)*8);
	for (int i = 0; i < 8; i++) {
		ss[i] = (int*)calloc(1, sizeof(int)*128);
	}
	int* ids = (int*)calloc(1, sizeof(int)*8*128);

	int row = 0, col = 0;

	int n = 0;
	char** lines = load_file("./input", 0, &n);
	char buf[128] = {0};

	int fr = 0, fc = 0;

	while(1) {
		cfx_clear();

		if (i < n) {
			storc(lines[i], &row, &col);
			ss[col][row] = 1;
			ids[(row<<3)+col] = 1;
			sprintf(buf, "%s -> (%d, %d)", lines[i], col, row);
		} else {
			if (fc == 0 && fr == 0) {
				get_mine(ids, &fc, &fr);
			}
			sprintf(buf, "Seat found -> (%d, %d) | id = %d", fc, fr, (fr<<3)+fc);
			int py = fc >= 4 ? 40 : 30;
			cfx_color(200, 50, 50);
			cfx_draw_rectangle(10+(fr*10), py+(fc*10), 10, 10, 1);
		}

		draw_seats(10, 30, ss);
		cfx_color(200, 200, 200);
		cfx_draw_text(100, 20, buf);

		// Update screen
		cfx_flush();

		i+=i<n;

		if (cfx_event_waiting()) {
			// Wait for the user to press a character.
			c = cfx_wait(&x,&y);
			printf("c=%c x=%d y=%d\n", c, x, y);

			// Quit if it is the letter q.
			if(c=='q') break;
		}

		// Sleep for enough time to update window at 60 fps
		cfx_wait_frame();
	}

	// Cleanup
	ffree((void**)ss, 8);
	ffree((void**)lines, n);
	free(ids);
	cfx_free();

	return 0;
}

