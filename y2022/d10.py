import aoc

import dataclasses
import enum
import logging

from queue import PriorityQueue
from typing import Any, List, Tuple, Dict, Optional


def part1(text: str) -> int:
    x = 1
    s = 0
    cycle = 1

    for line in text.splitlines():
        toks = line.split()
        if toks[0] == "addx":
            cycle += 1
            if (cycle-20) % 40 == 0:
                s += x*cycle
            cycle += 1
            x += int(toks[1])
            if (cycle-20) % 40 == 0:
                s += x*cycle
        else:
            cycle += 1
            if (cycle-20) % 40 == 0:
                s += x*cycle
    return s

def part2(text: str) -> str:
    x = 1
    cycle = 0

    p = 0
    out = "\n  "

    light = "█"
    dark = " "

    for line in text.splitlines():
        toks = line.split()
        if toks[0] == "addx":
            if (x-1) <= p <= (x+1):
                out += light
            else:
                out += dark
            p += 1

            cycle += 1
            if cycle % 40 == 0:
                out += "\n  "
                p = 0

            if (x-1) <= p <= (x+1):
                out += light
            else:
                out += dark
            p += 1

            cycle += 1
            if cycle % 40 == 0:
                out += "\n  "
                p = 0

            x += int(toks[1])
        else:
            if (x-1) <= p <= (x+1):
                out += light
            else:
                out += dark
            p += 1

            cycle += 1
            if cycle % 40 == 0:
                out += "\n  "
                p = 0

    return out

def main():
    d = 10
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
noop
addx 3
addx -5
""",
"""\
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
"""
        ]
        for test in tests:
            logging.info(f'test: "{None}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

