#!/usr/bin/env python
import argparse
from tqdm import tqdm

class CupGame:
    def __init__(self, line, max_val=1000000):
        self.input = [int(x) for x in line.strip()]
        self.max = max_val
        self.cups = [0 for i in range(self.max+1)]
        start = self.input[0]
        for i, c in enumerate(self.input):
            self.cups[c] = self.input[(i+1)%len(self.input)]
            prev = c
        for i in range(len(self.input)+1, self.max+1):
            self.cups[prev] = i
            prev = i
        self.cups[prev] = start
        self.current = start

    def step(self):
        c1 = self.cups[self.current]
        c2 = self.cups[c1]
        c3 = self.cups[c2]
        after = self.cups[c3]

        dest = self.current-1
        if dest == 0: dest = self.max
        while dest == c1 or dest == c2 or dest == c3:
            dest -= 1
            if dest == 0: dest = self.max

        self.cups[self.current] = after  # 'remove' next 3

        after_dest = self.cups[dest]     # save what follows dest
        self.cups[dest] = c1             # move next 3 to follow dest
        self.cups[c3] = after_dest       # complete insert after dest
        self.current = after             # set current for next iteration

    def simulate(self, n=10000000):
        for i in tqdm(range(n)):
            self.step()
        a = self.cups[1]
        b = self.cups[a]
        return a,b

def main(fn):
    with open(fn, 'r') as fp:
        cg = CupGame(next(fp))
        res = cg.simulate()
        print(f'result = {res}, {res[0]*res[1]}')


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)
