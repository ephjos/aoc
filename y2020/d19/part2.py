#!/usr/bin/env python
import argparse
import regex
from tqdm import tqdm

class RuleSet():
    def __init__(self, lines_dict):
        m = max([int(k) for k in lines_dict.keys()])
        lines = ["" for i in range(m+1)]
        for k,v in lines_dict.items():
            lines[int(k)] = v.strip().replace('\"', '')
        nums_left = True
        md = 1<<6
        d = 0
        while nums_left and d < md:
            nums_left = False
            for i, line in enumerate(lines):
                nums = set(regex.findall(r'\d+', line))
                nums_left |= nums != set()
                print(line)
                d += 1
                for num in nums:
                    to_sub = lines[int(num)]
                    has_or = '|' in to_sub
                    search = regex.compile(r'\b{}\b'.format(num))
                    if has_or:
                        lines[i] = regex.sub(
                            search, '({})'.format(to_sub), lines[i])
                    else:
                        lines[i] = regex.sub(
                            search, to_sub, lines[i])
                break

        self.rules = []
        for line in [l.replace(' ', '') for l in lines]:
            self.rules.append(line)

    def __getitem__(self, idx):
        return self.rules[idx]

def main(fn):
    with open(fn, 'r') as fp:
        rule_lines = {}
        get_rules = True
        messages = []
        count = 0
        ruleset = None
        for line in [l.strip() for l in fp]:
            if get_rules and line != "":
                idx, v = [tok.strip() for tok in line.split(':')]
                rule_lines[idx] = v
            elif line == "":
                ruleset = RuleSet(rule_lines)
                get_rules = False
            else:
                count += regex.fullmatch(ruleset[0], line) is not None

        print('count = {}'.format(count))



if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

