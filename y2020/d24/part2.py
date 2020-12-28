#!/usr/bin/env python
import argparse
from collections import defaultdict

# Combination of
# https://github.com/r0f1/adventofcode2020/blob/master/day24/main.py
#     and
# https://gist.github.com/aledesole/ca32d5cae663cb15a7be213542d52093
#
# Clean parsing from first link and general approach from second, improved by
# using a set of the black tiles instead of operating on a window

def parse(s):
    res = []
    while s:
        if   s[0]   == 'w': res.append((-1,0));  s = s[1:]
        elif s[0]   == 'e': res.append((1,0));   s = s[1:]
        elif s[0:2] == 'se': res.append((1,-1)); s = s[2:]
        elif s[0:2] == 'sw': res.append((0,-1)); s = s[2:]
        elif s[0:2] == 'ne': res.append((0,1));  s = s[2:]
        elif s[0:2] == 'nw': res.append((-1,1)); s = s[2:]
    return res

def apply(n, ns):
    if n == 0:
        if sum(ns) == 2:
            return 1
    else:
        if sum(ns) == 0 or sum(ns) > 2:
            return 0
    return n

def simulate(d, n=100):
    dirs = [(-1,0),(1,0),(1,-1),(0,-1),(0,1),(-1,1)]
    blacks = set(k for k,v in d.items() if v == 1)
    for _ in range(n):
        new_blacks = set()
        coords = set((y+i, x+j) for y,x in blacks
                  for i in range(-1,2) for j in range(-1,2))
        for coord in coords:
            i,j = coord
            if apply(int(coord in blacks),
                     [int(k in blacks) for k in [(y+i,x+j) for x,y in dirs]]):
                new_blacks.add(coord)
        blacks = new_blacks
    return blacks

def main(fn):
    with open(fn, 'r') as fp:
        d = defaultdict(int)
        lines = [parse(l.strip()) for l in fp]
        for line in lines:
            xs, ys = zip(*line)
            coords = sum(xs), sum(ys)
            d[coords] = 1 - d[coords]
        print(sum(d.values()))

        res_d = simulate(d.copy())
        print(len(res_d))


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

