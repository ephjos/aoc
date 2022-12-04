import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional

def part1(text: str) -> int:
    ranges = [set(list(range(int(range_raw.split("-")[0]), int(range_raw.split("-")[1])+1))) for line in text.splitlines() for range_raw in line.split(",")]
    s = 0
    for i in range(0, len(ranges), 2):
        s += ranges[i].issuperset(ranges[i+1]) or ranges[i+1].issuperset(ranges[i])

    return s

def part2(text: str) -> int:
    ranges = [set(list(range(int(range_raw.split("-")[0]), int(range_raw.split("-")[1])+1))) for line in text.splitlines() for range_raw in line.split(",")]
    s = 0
    for i in range(0, len(ranges), 2):
        s += not ranges[i].isdisjoint(ranges[i+1])

    return s

def main():
    d = 4
    text = aoc.get_input(2022, d).strip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.strip()}"\n  {d}.1: {part1(test.strip())}\n  {d}.2: {part2(test.strip())}')

if __name__ == "__main__":
    main()

