import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional


def part1(text: str) -> int:
    s = 0

    d = {}
    for i, v in enumerate(range(65, 91)):
        d[chr(v)] = i + 27
    for i, v in enumerate(range(97,123)):
        d[chr(v)] = i + 1

    for line in text.splitlines():
        l = len(line)
        h = l // 2
        seen = set(line[0:h]) & set(line[h:l])
        s += d[list(seen)[0]]

    return s

def part2(text: str) -> int:
    s = 0

    d = {}
    for i, v in enumerate(range(65, 91)):
        d[chr(v)] = i + 27
    for i, v in enumerate(range(97,123)):
        d[chr(v)] = i + 1

    lines = text.splitlines()
    for i in range(0, len(lines), 3):
        seen = set(lines[i]) & set(lines[i+1]) & set(lines[i+2])
        s += d[list(seen)[0]]

    return s

def main():
    d = 3
    text = aoc.get_input(2022, d).strip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.strip()}"\n  {d}.1: {part1(test.strip())}\n  {d}.2: {part2(test.strip())}')

if __name__ == "__main__":
    main()

