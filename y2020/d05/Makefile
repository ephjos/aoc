
CC = gcc -g
objects = part1 part2

.PHONY: all
all: $(objects)

clean:
	rm -rfv $(objects)

$(objects): %: %.c
	$(CC) -o $@ $<

.PHONY: vis
vis:
	$(CC) -lX11 -lm -lavdevice -lavformat -lavfilter -lavcodec -lswresample \
		-lswscale -lavutil -o vis vis.c




