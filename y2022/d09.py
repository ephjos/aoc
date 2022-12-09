import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional

def sign(a):
    return (a > 0) - (a < 0)

def part1(text: str) -> int:
    # i, j
    h = [0, 0]
    t = [0, 0]
    t_seen = {
        tuple(t): 1,
    }

    for line in text.splitlines():
        card, steps = line.split()
        for _ in range(int(steps)):
            if card == 'R':
                h[1] += 1
            elif card == 'L':
                h[1] -= 1
            elif card == 'U':
                h[0] -= 1
            elif card == 'D':
                h[0] += 1
            else:
                raise Exception(f'Unknown cardinal "{card}"')


            di = h[0]-t[0]
            dj = h[1]-t[1]
            if di == 0 and abs(dj) == 2:
                t[1] += sign(dj)
            elif dj == 0 and abs(di) == 2:
                t[0] += sign(di)
            elif abs(di) > 1 or abs(dj) > 1:
                t[0] += sign(di)
                t[1] += sign(dj)


            t_seen[tuple(t)] = t_seen.get(tuple(t), 0)+1

    return len(list(t_seen.keys()))

def part2(text: str) -> int:
    rope = [[0, 0] for _ in range(10)]
    seen = [{} for _ in range(10)]

    for line in text.splitlines():
        card, steps = line.split()
        for _ in range(int(steps)):
            if card == 'R':
                rope[0][1] += 1
            elif card == 'L':
                rope[0][1] -= 1
            elif card == 'U':
                rope[0][0] -= 1
            elif card == 'D':
                rope[0][0] += 1
            else:
                raise Exception(f'Unknown cardinal "{card}"')

            for i in range(9):
                h = rope[i]
                t = rope[i+1]

                di = h[0]-t[0]
                dj = h[1]-t[1]
                if di == 0 and abs(dj) == 2:
                    t[1] += sign(dj)
                elif dj == 0 and abs(di) == 2:
                    t[0] += sign(di)
                elif abs(di) > 1 or abs(dj) > 1:
                    t[0] += sign(di)
                    t[1] += sign(dj)

                seen[i+1][tuple(t)] = seen[i+1].get(tuple(t), 0)+1

    return len(list(seen[9].keys()))

def main():
    d = 9
    text = aoc.get_input(2022, d).rstrip()
    print(text) # TODO: remove

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
""",
"""\
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

