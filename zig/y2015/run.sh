#!/bin/sh

find . | entr -c -s 'zig build run'
