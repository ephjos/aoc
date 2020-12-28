#!/usr/bin/env python
import argparse

def get_loop_size(key, modulo):
    sub = 7
    val = 1
    loop_size = 0
    while val != key:
        val *= sub
        val %= modulo
        loop_size += 1
    return loop_size

def parse_keys(fn):
    with open(fn, 'r') as fp:
        pub_keys = []
        for line in [l.strip() for l in fp]:
            pub_keys.append(int(line))
        return pub_keys

def main(fn):
    modulo = 20201227
    pub_keys = parse_keys(fn)
    loop_sizes = [get_loop_size(pub_key, modulo) for pub_key in pub_keys]

    val = 1
    for _ in range(loop_sizes[0]):
        val *= pub_keys[1]
        val %= modulo
    print(val)

    val = 1
    for _ in range(loop_sizes[1]):
        val *= pub_keys[0]
        val %= modulo
    print(val)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

