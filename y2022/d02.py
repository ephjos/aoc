import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional


def part1(text: str) -> int:
    s = 0
    d = {
        'A': 1,
        'X': 1,
        'B': 2,
        'Y': 2,
        'C': 3,
        'Z': 3,
    }

    for line in text.split("\n"):
        o, u = line.split(" ")
        s += d[u]
        if d[u] == d[o] + 1 or (d[u] == 1 and d[o] == 3):
            s += 6
        elif d[u] == d[o]:
            s += 3

    return s

def part2(text: str) -> int:
    s = 0
    d = {
        'A': 1,
        'X': 0,
        'B': 2,
        'Y': 3,
        'C': 3,
        'Z': 6,
    }

    for line in text.split("\n"):
        o, need = line.split(" ")
        s += d[need]

        if d[need] == 0:
            if d[o] == 1:
                s += 3
            else:
                s += d[o] - 1
        elif d[need] == 3:
            s += d[o]
        else:
            if d[o] == 3:
                s += 1
            else:
                s += d[o] + 1

    return s

def main():
    d = 2
    text = aoc.get_input(2022, d).strip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
            """
A Y
B X
C Z
            """
        ]
        for test in tests:
            logging.info(f'test: "{test.strip()}"\n  {d}.1: {part1(test.strip())}\n  {d}.2: {part2(test.strip())}')

if __name__ == "__main__":
    main()

