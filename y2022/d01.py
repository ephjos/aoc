import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional


def part1(text):
    elves = text.split("\n\n")
    m = float('-inf')
    for e in elves:
        s = 0
        for n in e.split("\n"):
            s += int(n)
        m = max(m, s)
    return m

def part2(text):
    elves = text.split("\n\n")
    s = [sum(map(int, e.split("\n"))) for e in elves]
    return sum(sorted(s, reverse=True)[0:3])

def main():
    d = 1
    text = aoc.get_input(2022, d).strip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
            """
            1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
            """,
        ]
        for test in tests:
            logging.info(f'test: "{test.strip()}"\n  {d}.1: {part1(test.strip())}\n  {d}.2: {part2(test.strip())}')

if __name__ == "__main__":
    main()

