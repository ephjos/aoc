import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional


def part1(text):
    return 0

def part2(text):
    return 0

def main():
    d = 1
    text = aoc.get_input(2022, d)

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
        ]
        for test in tests:
            logging.info(f'test: "{test}"\n  {d}.1: {part1(test)}\n  {d}.2: {part2(test)}')

if __name__ == "__main__":
    main()

