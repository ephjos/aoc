import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional

ROCK = 1
SAND = 2
SOURCE = 3

CHARS = {
    1: "#",
    2: "o",
    3: "+",
}

def print_point_dict(d: Dict[Tuple[int,int], int]) -> None:
    min_x = float("infinity")
    min_y = float("infinity")
    max_x = float("-infinity")
    max_y = float("-infinity")
    for p in d.keys():
        min_x = min(min_x, p[0])
        min_y = min(min_y, p[1])
        max_x = max(max_x, p[0])
        max_y = max(max_y, p[1])

    min_x -= 200
    min_y -= 5
    max_x += 200
    max_y += 5


    out = ""
    for i in range(min_y, max_y+1):
        for j in range(min_x, max_x+1):
            p = (j, i)
            if p in d:
                out += CHARS[d[p]]
            else:
                out += "."
        out += "\n"
    print(out)

def sand_move(d, s) -> Optional[Tuple[int, int]]:
    below = (s[0], s[1]+1)
    if below not in d:
        return below
    below_left= (s[0]-1, s[1]+1)
    if below_left not in d:
        return below_left
    below_right= (s[0]+1, s[1]+1)
    if below_right not in d:
        return below_right
    return None

def part1(text: str) -> int:
    d = {}
    max_y = float("-infinity")
    for line in text.splitlines():
        points = list(map(lambda x: tuple(map(int, x.strip().split(","))), line.split("->")))
        for i in range(len(points)-1):
            a = points[i]
            b = points[i+1]

            dx = b[0] - a[0]
            dy = b[1] - a[1]

            for j in range(min(a[0], a[0]+dx), max(a[0], a[0]+dx)+1):
                p = (j, a[1])
                d[p] = 1
                max_y = max(max_y, a[1])

            for j in range(min(a[1], a[1]+dy), max(a[1], a[1]+dy)+1):
                p = (a[0], j)
                d[p] = 1
                max_y = max(max_y, j)

    d[(500,0)] = SOURCE

    sand_filling = True
    sand_count = 0
    while True:
        sand = (500, 0)
        while move := sand_move(d, sand):
            sand = move
            if sand[1] > max_y:
                sand_filling = False
                break

        if not sand_filling:
            break
        d[sand] = SAND
        sand_count += 1
    return sand_count

def sand_move_2(d, s, max_y) -> Optional[Tuple[int, int]]:
    below = (s[0], s[1]+1)
    if below not in d and below[1] < max_y:
        return below
    below_left= (s[0]-1, s[1]+1)
    if below_left not in d and below_left[1] < max_y:
        return below_left
    below_right= (s[0]+1, s[1]+1)
    if below_right not in d and below_right[1] < max_y:
        return below_right
    return None

def part2(text: str) -> int:
    d = {}
    max_y = float("-infinity")
    for line in text.splitlines():
        points = list(map(lambda x: tuple(map(int, x.strip().split(","))), line.split("->")))
        for i in range(len(points)-1):
            a = points[i]
            b = points[i+1]

            dx = b[0] - a[0]
            dy = b[1] - a[1]

            for j in range(min(a[0], a[0]+dx), max(a[0], a[0]+dx)+1):
                p = (j, a[1])
                d[p] = 1
                max_y = max(max_y, a[1])

            for j in range(min(a[1], a[1]+dy), max(a[1], a[1]+dy)+1):
                p = (a[0], j)
                d[p] = 1
                max_y = max(max_y, j)

    d[(500,0)] = SOURCE
    max_y += 2

    sand_filling = True
    sand_count = 0
    while True:
        sand = (500, 0)
        while move := sand_move_2(d, sand, max_y):
            sand = move

        d[sand] = SAND
        sand_count += 1

        if sand == (500, 0):
            break
    return sand_count

def main():
    d = 14
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

