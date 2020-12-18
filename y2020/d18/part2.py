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

def insert_parens(line):
    i = 0
    while i < len(line):
        c = line[i]
        if c == '+':
            # insert left
            j = i-1
            c = 0
            f = 0
            while j > 0:
                c += line[j] == ')'
                c -= line[j] == '('
                if line[j] != ' ' and line[j] != ')':
                    if c == 0:
                        line.insert(j, '(')
                        i += 1
                        f = 1
                        break
                j -= 1
            if not f:
                line.insert(0, '(')
                i += 1
            # insert right
            j = i+1
            c = 1
            while j < len(line):
                c += line[j] == '('
                c -= line[j] == ')'
                if line[j] != ' ' and line[j] != '(':
                    if c == 1:
                        line.insert(j+1, ')')
                        i += 2
                        break
                j += 1
        #print(''.join(line))
        i += 1
    return line

def main(fn):
    with open(fn, 'r') as fp:
        s = 0
        for line in fp:
            line = list(line.strip())
            line = insert_parens(line)
            i, v = parse(line)
            print('{} = {}'.format(''.join(line),  v))
            s += v
        print('sum = {}'.format(s))

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

