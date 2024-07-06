#!/bin/sh

zig build -Dbenchmark=true -Doptimize=ReleaseFast && ./zig-out/bin/y2015

# TODO: gnuplot

