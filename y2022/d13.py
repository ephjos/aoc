import aoc

import ast
import dataclasses
import enum
import functools
import logging

from typing import Any, List, Tuple, Dict, Optional

def parse_packet(line: str) -> List:
    return ast.literal_eval(line)


def compare(left, right) -> int:
    """
    return -1 if left < right
    return 0 if left == right
    return 1 if left > right
    """
    for i in range(max(len(left), len(right))):
        try:
            l = left[i]
        except IndexError:
            return -1

        try:
            r = right[i]
        except IndexError:
            return 1

        l_is_int = isinstance(l, int)
        r_is_int = isinstance(r, int)
        if l_is_int and r_is_int:
            if l < r:
                return -1
            elif l > r:
                return 1
        elif l_is_int:
            c = compare([l], r)
            if c == -1 or c == 1:
                return c
        elif r_is_int:
            c = compare(l, [r])
            if c == -1 or c == 1:
                return c
        else:
            c = compare(l, r)
            if c == -1 or c == 1:
                return c

    return 0

def part1(text: str) -> int:
    packet_groups = [[parse_packet(packet) for packet in block.splitlines()] for block in text.split("\n\n")]
    sum_of_indices = 0
    for i, packet_group in enumerate(packet_groups, 1):
        left, right = packet_group
        if compare(left, right) == -1:
            sum_of_indices += i
    return sum_of_indices

def part2(text: str) -> int:
    packets = [parse_packet("[[2]]"), parse_packet("[[6]]")]
    for line in text.splitlines():
        if line.strip() == "":
            continue
        packets.append(parse_packet(line))

    sorted_packets = sorted(packets, key=functools.cmp_to_key(compare))

    mul_of_indices = 1
    for i, p in enumerate(sorted_packets, 1):
        if p == packets[0] or p == packets[1]:
            mul_of_indices *= i
    return mul_of_indices

def main():
    d = 13
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
""",
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

