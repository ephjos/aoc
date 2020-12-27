#!/usr/bin/env python
import argparse

class CupGame:
    def __init__(self, line):
        self.input = line.strip()
        self.cups = [0 for i in range(10)]
        for i, c in enumerate(self.input):
            self.cups[int(c)] = int(self.input[(i+1)%9])
        self.current = int(self.input[0])

    def step(self):
        c1 = self.cups[self.current]
        c2 = self.cups[c1]
        c3 = self.cups[c2]
        after = self.cups[c3]

        dest = self.current-1
        if dest == 0: dest = 9
        while dest == c1 or dest == c2 or dest == c3:
            dest -= 1
            if dest == 0: dest = 9

        self.cups[self.current] = after  # 'remove' next 3

        after_dest = self.cups[dest]     # save what follows dest
        self.cups[dest] = c1             # move next 3 to follow dest
        self.cups[c3] = after_dest       # complete insert after dest
        self.current = after             # set current for next iteration

    def simulate(self, n=100):
        for i in range(n):
            self.step()
        return str(self)

    def __str__(self):
        out = ''
        curr = self.cups[1]
        while curr != 1:
            out += str(curr)
            curr = self.cups[curr]
        return out

def main(fn):
    with open(fn, 'r') as fp:
        cg = CupGame(next(fp))
        res = cg.simulate(100)
        print(f'result = {res}')


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)
