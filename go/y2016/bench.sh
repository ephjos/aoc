#!/bin/bash

set -euo pipefail

go test -bench=. -benchmem -benchtime=10x | tee /tmp/aoc_go_2016_bench

# Plot
head -n -2 /tmp/aoc_go_2016_bench | tail -n+5 | sed -e 's/BenchmarkAll\/day\(.*\)-6/\1/' | awk -F " " '{ x+=$3; print NR" "$1" "$3/1000000" "$5" "$7" "x/1000000 }' > /tmp/aoc_go_2016_bench.dat
gnuplot -p bench.gp 

# If formatting breaks, just use this
head -n -2 /tmp/aoc_go_2016_bench | tail -n+4 | awk -F " " '{ runs+=$2; ns+=$3; b+=$5; allocs+=$7; } END { printf "\n%s\n\t%d\n\t%d ms/run\n\t%d KiB/run\n\t%d allocs/run\n", "Total:", runs, ns/1000000, b/1024, allocs }'
