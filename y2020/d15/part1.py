#!/usr/bin/env python

def play_game(line, to):
    d = {}
    i = 0
    prev = None
    for n in [int(x) for x in line.split(",")]:
        d[n] = i
        i += 1
        prev = n

    while i < to:
        if prev not in d:
            d[prev] = i-1
        t = d[prev]
        x = i - t -1
        d[prev] = i-1
        prev = x
        i += 1
    print('{}th word: {}'.format(to, prev))

def main():
    with open('./test') as fp:
        for line in fp:
            play_game(line, 2020)

if __name__ == "__main__":
    main()

