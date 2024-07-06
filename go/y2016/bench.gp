
#set terminal dumb
set terminal png

set grid
set logscale y

set xtics (\
"01_a" 1,\
"01_b" 2,\
"02_a" 3,\
"02_b" 4,\
"03_a" 5,\
"03_b" 6,\
"04_a" 7,\
"04_b" 8,\
"05_a" 9,\
"05_b" 10,\
"06_a" 11,\
"06_b" 12,\
"07_a" 13,\
"07_b" 14,\
"08_a" 15,\
"08_b" 16,\
"09_a" 17,\
"09_b" 18,\
"10_a" 19,\
"10_b" 20,\
"11_a" 21,\
"11_b" 22,\
"12_a" 23,\
"12_b" 24,\
"13_a" 25,\
"13_b" 26,\
"14_a" 27,\
"14_b" 28,\
"15_a" 29,\
"15_b" 30,\
"16_a" 31,\
"16_b" 32,\
"17_a" 33,\
"17_b" 30,\
"18_a" 35,\
"18_b" 36,\
"19_a" 37,\
"19_b" 38,\
"20_a" 39,\
"20_b" 40,\
"21_a" 41,\
"21_b" 42,\
"22_a" 43,\
"22_b" 44,\
"23_a" 45,\
"23_b" 46,\
"24_a" 47,\
"24_b" 48,\
"25_a" 49,\
"25_b" 50,\
)

set output 'images/ms.png'
set title "Runtime (ms) per part"
plot '/tmp/aoc_go_2016_bench.dat' using 1:3 with linespoints linewidth 3 linecolor rgb "red" notitle

set output 'images/b.png'
set title "Bytes allocated per part"
plot '/tmp/aoc_go_2016_bench.dat' using 1:4 with linespoints linewidth 3 linecolor rgb "green" notitle

set output 'images/allocs.png'
set title "Allocations per part"
plot '/tmp/aoc_go_2016_bench.dat' using 1:5 with linespoints linewidth 3 linecolor rgb "blue" notitle

set output 'images/total.png'
set title "Cumulative Runtime (ms)"
plot '/tmp/aoc_go_2016_bench.dat' using 1:6 with linespoints linewidth 3 linecolor rgb "purple" notitle
