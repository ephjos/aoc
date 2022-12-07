import aoc

import dataclasses
import enum
import pathlib
import logging

from typing import Any, List, Tuple, Dict, Optional


def get_dir_size(d):
    return 0

def part1(text: str) -> int:
    cd = []
    d = {}

    lines = text.splitlines()
    i = 0

    while i < len(lines):
        toks = lines[i].split()
        if toks[1] == "cd":
            if toks[2] == "..":
                cd.pop()
            else:
                cd.append(toks[2])
                t = d
                for x in cd:
                    t[x] = t.get(x, {})
                    t = t[x]

            i += 1
        elif toks[1] == "ls":
            i += 1
            t = d
            for x in cd:
                t = t[x]
            while i < len(lines) and lines[i][0] != "$":
                toks = lines[i].split()
                if toks[0] != "dir":
                    t[toks[1]] = int(toks[0])
                i += 1

    dir_sizes = {}

    def walk(ik, d, indent=0):
        if isinstance(d, int):
            return d

        s = 0
        for k in d.keys():
            #print((' '*indent) + k)
            s += walk(ik + "/" + k, d[k], indent+2)
        dir_sizes[ik] = s
        return s

    walk('/', d)

    return sum(filter(lambda x: x <= 100000 ,dir_sizes.values()))

def part2(text: str) -> int:
    cd = []
    d = {}

    lines = text.splitlines()
    i = 0

    while i < len(lines):
        toks = lines[i].split()
        if toks[1] == "cd":
            if toks[2] == "..":
                cd.pop()
            else:
                cd.append(toks[2])
                t = d
                for x in cd:
                    t[x] = t.get(x, {})
                    t = t[x]

            i += 1
        elif toks[1] == "ls":
            i += 1
            t = d
            for x in cd:
                t = t[x]
            while i < len(lines) and lines[i][0] != "$":
                toks = lines[i].split()
                if toks[0] != "dir":
                    t[toks[1]] = int(toks[0])
                i += 1

    dir_sizes = {}

    def walk(ik, d, indent=0):
        if isinstance(d, int):
            return d

        s = 0
        for k in d.keys():
            #print((' '*indent) + k)
            s += walk(ik + "/" + k, d[k], indent+2)
        dir_sizes[ik] = s
        return s

    walk('/', d)

    TOTAL_SPACE = 70000000
    NEEDED_SPACE = 30000000
    need_to_delete = NEEDED_SPACE - (TOTAL_SPACE - dir_sizes["/"])

    return list(filter(lambda x: x >= need_to_delete, sorted(dir_sizes.values())))[0]

def main():
    d = 7
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

