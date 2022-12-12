import aoc

import dataclasses
import enum
import logging
from queue import Queue, PriorityQueue
import heapq

from typing import Any, List, Tuple, Dict, Optional


def part1(text: str) -> int:
    DIRS = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    heightmap = []
    S = ()
    E = ()


    for i, line in enumerate(text.splitlines()):
        heightmap_row = [ord(c)-97 for c in line]
        if 'S' in line:
            j = line.index('S')
            heightmap_row[j] = 0
            S = (i, j)
        if 'E' in line:
            j = line.index('E')
            heightmap_row[j] = 25
            E = (i, j)
        heightmap.append(heightmap_row)

    h = len(heightmap)
    w = len(heightmap[0])

    dist = dict([((i,j), float("infinity")) for i in range(h) for j in range(w)])
    dist[S] = 0
    prev = {}
    Q = [(0, S)]

    while len(Q) > 0:
        cd, u = heapq.heappop(Q)

        neighbors = filter(lambda n: n[0] >= 0 and n[0] < h and n[1] >= 0 and n[1] < w and (heightmap[n[0]][n[1]] - heightmap[u[0]][u[1]]) <= 1, [(u[0] + d[0], u[1] + d[1]) for d in DIRS])
        for v in neighbors:
            alt = dist[u] + 1
            if alt < dist[v]:
                dist[v] = alt
                prev[v] = u
                heapq.heappush(Q, (alt, v))

#    prevs = []
#    c = E
#    while c in prev:
#        prevs.append(c)
#        c = prev[c]
#
#    out = ""
#    for i in range(h):
#        for j in range(w):
#            if (i,j) == E:
#                out += "E"
#            elif (i,j) == S:
#                out += "s"
#            else:
#                out += chr(heightmap[i][j]+97) if (i,j) in prevs else "."
#        out += "\n"
#    print(out)

    return dist[E]

def part2(text: str) -> int:
    DIRS = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    heightmap = []
    Ss = []
    E = ()


    for i, line in enumerate(text.splitlines()):
        heightmap_row = [ord(c)-97 for c in line]
        for j, c in enumerate(line):
            if c == 'a':
                Ss.append((i,j))
        if 'S' in line:
            j = line.index('S')
            heightmap_row[j] = 0
            Ss.append((i,j))
        if 'E' in line:
            j = line.index('E')
            heightmap_row[j] = 25
            E = (i, j)
        heightmap.append(heightmap_row)

    def run(S):
        h = len(heightmap)
        w = len(heightmap[0])

        dist = dict([((i,j), float("infinity")) for i in range(h) for j in range(w)])
        dist[S] = 0
        prev = {}
        Q = [(0, S)]

        while len(Q) > 0:
            cd, u = heapq.heappop(Q)

            neighbors = filter(lambda n: n[0] >= 0 and n[0] < h and n[1] >= 0 and n[1] < w and (heightmap[n[0]][n[1]] - heightmap[u[0]][u[1]]) <= 1, [(u[0] + d[0], u[1] + d[1]) for d in DIRS])
            for v in neighbors:
                alt = dist[u] + 1
                if alt < dist[v]:
                    dist[v] = alt
                    prev[v] = u
                    heapq.heappush(Q, (alt, v))

#        prevs = []
#        c = E
#        while c in prev:
#            prevs.append(c)
#            c = prev[c]
#
#        out = ""
#        for i in range(h):
#            for j in range(w):
#                if (i,j) == E:
#                    out += "E"
#                elif (i,j) == S:
#                    out += "s"
#                else:
#                    out += chr(heightmap[i][j]+97) if (i,j) in prevs else "."
#            out += "\n"
#        print(out)

        return dist[E]

    return min([run(S) for S in Ss])

def main():
    d = 12
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

