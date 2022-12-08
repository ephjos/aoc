import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional


def part1(text: str) -> int:
    g = []
    for line in text.splitlines():
        g.append([int(c) for c in line])

    w = len(g[0])
    h = len(g)

    # Length of border is perimeter - 4 since the corners overlap
    border = (2*w) + (2*h) - 4

    def is_visible(i, j, g):

        pred = lambda x: x >= g[i][j]

        above = [g[ni][j] for ni in range(0, i)]
        if len(list(filter(pred, above))) == 0:
            return True

        below = [g[ni][j] for ni in range(i+1, h)]
        if len(list(filter(pred, below))) == 0:
            return True

        left = [g[i][nj] for nj in range(0, j)]
        if len(list(filter(pred, left))) == 0:
            return True

        right = [g[i][nj] for nj in range(j+1, w)]
        if len(list(filter(pred, right))) == 0:
            return True

        return False

    interior_visible = 0
    for i in range(1, h-1):
        for j in range(1, w-1):
            interior_visible += is_visible(i, j, g)

    return border + interior_visible

def part2(text: str) -> int:
    g = []
    for line in text.splitlines():
        g.append([int(c) for c in line])

    w = len(g[0])
    h = len(g)

    def scenic_score(i, j, g):
        above_score = 0
        below_score = 0
        left_score = 0
        right_score = 0

        curr = g[i][j]

        above = [g[ni][j] for ni in range(i-1, -1, -1)]
        for x in above:
            above_score += 1
            if x >= curr:
                break

        below = [g[ni][j] for ni in range(i+1, h, 1)]
        for x in below:
            below_score += 1
            if x >= curr:
                break

        left = [g[i][nj] for nj in range(j-1, -1, -1)]
        for x in left:
            left_score += 1
            if x >= curr:
                break

        right = [g[i][nj] for nj in range(j+1, w, 1)]
        for x in right:
            right_score += 1
            if x >= curr:
                break

        return above_score * below_score * left_score * right_score

    m = -1
    for i in range(0, h):
        for j in range(0, w):
            m = max(m, scenic_score(i, j, g))

    return m

def main():
    d = 8
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
30373
25512
65332
33549
35390
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

