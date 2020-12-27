#!/usr/bin/env python
import argparse

def walk(tiles, dirs):
    curr = (0,0)

    i = 0
    l = len(dirs)
    while i < l:
        c = dirs[i]
        inc = 1
        if c == 'e': # east
            curr = (curr[0], curr[1]+2)
        elif c == 'w': # west
            curr = (curr[0], curr[1]-2)
        elif c == 's':
            if dirs[i+1] == 'e': # south east
                curr = (curr[0]+1, curr[1]+1)
            elif dirs[i+1] == 'w': # south west
                curr = (curr[0]+1, curr[1]-1)
            inc = 2
        elif c == 'n':
            if dirs[i+1] == 'e': # north east
                curr = (curr[0]-1, curr[1]+1)
            elif dirs[i+1] == 'w': # north west
                curr = (curr[0]-1, curr[1]-1)
            inc = 2
        i += inc
    tiles[curr] = not tiles.get(curr, True)
    return tiles

def main(fn):
    with open(fn, 'r') as fp:
        tiles = {}
        for line in [l.strip() for l in fp]:
            walk(tiles, line)
        num_black = len(list(filter(lambda x: not x, tiles.values())))
        print(f'num black = {num_black}')


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

