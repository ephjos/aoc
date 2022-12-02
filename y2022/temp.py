import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional


def part1(text: str) -> int:
    return 0

def part2(text: str) -> int:
    return 0

def main():
    d = 1
    text = aoc.get_input(2022, d).strip()
    print(text) # TODO: remove

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
        ]
        for test in tests:
            logging.info(f'test: "{test.strip()}"\n  {d}.1: {part1(test.strip())}\n  {d}.2: {part2(test.strip())}')

if __name__ == "__main__":
    main()

