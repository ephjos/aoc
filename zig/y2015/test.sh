#!/bin/sh

find . | entr -c -s 'zig build test'
