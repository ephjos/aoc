#!/usr/bin/env python
import argparse
from copy import deepcopy
import math
from tqdm import tqdm

TILE_N = 10

class Side():
    TOP = 0
    RIGHT = 1
    BOTTOM = 2
    LEFT = 3

def pl(ls):
    for l in ls:
        print(l)
    print()

class Tile():
    def __init__(self, input_id, lines, build_orientations=True):
        self.id = input_id
        self.lines = [list(row) for row in lines]
        if build_orientations:
            self.orientations = self.build_orientations()

    def build_orientations(self):
        def flip_x(lines):
            res = [list(reversed(row)) for row in lines]
            return res

        def flip_y(lines):
            return list(reversed(deepcopy(lines)))

        def rotate(lines, n=1):
            """Rotates 90 deg clockwise"""
            if n < 1: return lines
            res = deepcopy(lines)
            for i in range(len(res)):
                for j in range(len(res)):
                    res[i][j] = lines[j][i]
            return rotate(flip_x(res), n-1)

        os = []

        for i in range(4):
            os.append(rotate(self.lines, i))
            os.append(flip_x(rotate(self.lines, i)))
            os.append(flip_y(rotate(self.lines, i)))
            os.append(flip_x(flip_y(rotate(self.lines, i))))

        return os

    def match(self, other):
        def top(lines): return lines[0]
        def bottom(lines): return lines[-1]
        def left(lines): return [r[0] for r in lines]
        def right(lines): return [r[-1] for r in lines]

        sides = []
        os = []

        for o in other.orientations:
            if top(self.lines) == bottom(o):
                sides.append(Side.TOP)
                os.append(o)
            if right(self.lines) == left(o):
                sides.append(Side.RIGHT)
                os.append(o)
            if bottom(self.lines) == top(o):
                sides.append(Side.BOTTOM)
                os.append(o)
            if left(self.lines) == right(o):
                sides.append(Side.LEFT)
                os.append(o)

        return sides, os

    def __str__(self):
        return str(self.id)

    def __repr__(self):
        return str(self)

class TileParser():
    def __init__(self, fp):
        self.fp = (l.strip() for l in fp)

    def parse(self):
        tiles = []
        for i, line in enumerate(self.fp):
            if line == '': line = next(self.fp)
            tile_id = int(line[:-1].split(' ')[-1])
            lines = [next(self.fp) for i in range(TILE_N)]
            tiles.append(Tile(tile_id, lines))

        return tiles

class ImageBuilder():
    def __init__(self, tiles):
        self.tiles = tiles
        self.n = math.isqrt(len(tiles))

    def build(self):
        tiles = deepcopy(self.tiles)
        board = {
            (0,0): tiles[0]
        }
        del tiles[0]

        pbar = tqdm(desc="Tiles left", total=len(tiles))
        while tiles:
            for k, bt in board.items():
                added = False
                for i, tile in enumerate(tiles):
                    sides, lls = bt.match(tile)
                    if not sides: continue
                    for s, ls in zip(sides, lls):
                        nt = Tile(tile.id, ls, build_orientations=False)
                        if s == Side.TOP:
                            nk = (k[0],k[1]+1)
                        elif s == Side.RIGHT:
                            nk = (k[0]+1,k[1])
                        elif s == Side.BOTTOM:
                            nk = (k[0],k[1]-1)
                        elif s == Side.LEFT:
                            nk = (k[0]-1,k[1])

                        if nk in board:
                            continue

                        board[nk] = nt
                        del tiles[i]
                        added = True
                        break
                    if added:
                        break
                if added:
                    pbar.update()
                    break

        pbar.close()
        out = [["" for i in range(self.n)] for i in range(self.n)]
        x,y = min(board.keys())
        for (u,v), t in sorted(board.items()):
            out[y-v-1][u-x] = str(t)
        pl(out)
        a, b = min(board.keys())
        c, d = max(board.keys())
        mul = board[(a,b)].id * board[(c,d)].id * \
            board[(a,d)].id * board[(c,b)].id
        print("mul = {}".format(mul))


def main(fn):
    ib = None
    with open(fn, 'r') as fp:
        tp = TileParser(fp)
        tiles = tp.parse()
        ib = ImageBuilder(tiles)
    ib.build()


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

