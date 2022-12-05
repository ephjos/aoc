import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional


def part1(text: str) -> str:
    sections = text.split("\n\n")
    starting_crates = list(reversed(sections[0].splitlines()))

    stacks = {}
    stack_names = starting_crates[0]
    for i in range(1, len(stack_names), 4):
        stacks[stack_names[i]] = []

    for line in starting_crates[1:]:
        for i in range(1, len(line), 4):
            if line[i] != " ":
                k = str(int(((i-1) / 4) + 1))
                stacks[k].append(line[i])

    instructions = sections[1]
    for line in instructions.splitlines():
        toks = line.split()
        n = int(toks[1])
        src = toks[3]
        dst = toks[5]
        for _ in range(n):
            stacks[dst].append(stacks[src].pop())

    out = ''.join([stack.pop() for stack in stacks.values()])
    return out

def part2(text: str) -> str:
    sections = text.split("\n\n")
    starting_crates = list(reversed(sections[0].splitlines()))

    stacks = {}
    stack_names = starting_crates[0]
    for i in range(1, len(stack_names), 4):
        stacks[stack_names[i]] = []

    for line in starting_crates[1:]:
        for i in range(1, len(line), 4):
            if line[i] != " ":
                k = str(int(((i-1) / 4) + 1))
                stacks[k].append(line[i])

    instructions = sections[1]
    for line in instructions.splitlines():
        toks = line.split()
        n = int(toks[1])
        src = toks[3]
        dst = toks[5]
        t = []
        for _ in range(n):
            t.append(stacks[src].pop())
        for v in reversed(t):
            stacks[dst].append(v)

    out = ''.join([stack.pop() for stack in stacks.values()])
    return out

def main():
    d = 5
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

