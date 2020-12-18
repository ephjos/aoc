#!/usr/bin/env python
import argparse

def parse_term(line, i):
    if line[i] == '(':
        return parse(line, i+1) # assume no space after (

    tok = line[i] # assume single digit followed by space
    return consume(line, i+1), int(tok)

ops = {
    '+': lambda x, y: x+y,
    '*': lambda x, y: x*y,
}
def parse_exprp(line, i, val):
    if i >= len(line):
        return i, val
    while i < len(line) and line[i] != ')':
        op = ops[line[i]]
        i, tv = parse_term(line, i+2)
        val = op(val, tv)
        i = consume(line, i)

    if i < len(line) and line[i] == ')':
        i += 1
    return consume(line, i), val

def consume(line, i):
    if i < len(line) and line[i] == ' ':
        return i+1
    return i

# expr  -> term expr'
# expr' -> ('+' | '*') term expr' | eps
# term  -> NUMBER | '(' expr ')'
def parse(line, i=0):
    i, val = parse_term(line, i)
    i, val = parse_exprp(line, i, val)
    return i, val

def main(fn):
    with open(fn, 'r') as fp:
        s = 0
        for line in fp:
            line = line.strip()
            i, v = parse(line)
            print('{} = {}'.format(line,  v))
            s += v
        print('sum = {}'.format(s))

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

