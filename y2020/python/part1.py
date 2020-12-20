#!/usr/bin/env python
import argparse

def main(fn):
    with open(fn, 'r') as fp:
        for line in [l.strip() for l in fp]:
            print(line)

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

