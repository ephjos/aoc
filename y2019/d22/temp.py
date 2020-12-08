#!/usr/bin/env python

def get_pos(x, p, lines):
    for line in lines:
        toks = line.strip().split(' ')
        op = ' '.join(toks[:2])
        if op == "deal with":
            x = (x * int(toks[3])) % p   # F(x) = x*n % p
        elif op == "deal into":
            x = p - x - 1                # F(x) = p-x-1
        else: # cut
            x = (x - int(toks[1])) % p   # F(x) = x - n % p
    return x

dw = lambda x, n, p: (x*n) % p
di = lambda x, n, p: p-x-1 % p
cu = lambda x, n, p: (x - n) % p

def main():
    lines = []
    with open("./input", "r") as fp:
        lines = fp.readlines()

    p = 10007
    print(get_pos(2019,p,lines))

if __name__ == "__main__":
    main()

