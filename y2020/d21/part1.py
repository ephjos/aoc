#!/usr/bin/env python
import argparse

def read_foods(fn):
    foods = []
    possible = {}
    allergen_is = {}
    with open(fn, 'r') as fp:
        for i, line in enumerate(l.strip() for l in fp):
            ings, alls = line.split(' (contains ')
            ings = set(ings.split(' '))
            alls = set(alls[:-1].split(', '))

            foods.append(ings)

            for al in alls:
                allergen_is[al] = allergen_is.get(al, [])
                allergen_is[al].append(i)

            for ing in ings:
                possible[ing] = possible.get(ing, set())
                possible[ing] |= alls

        return foods, possible, allergen_is

def get_safe(foods, possible, allergen_is):
    safe = []
    for ing, poss in possible.items():
        imposs = set()
        for al in poss:
            if any(ing not in foods[i] for i in allergen_is[al]):
                imposs.add(al)

        poss -= imposs
        if not poss:
            safe.append(ing)

    return safe

def main(fn):
    foods, possible, allergen_is = read_foods(fn)
    safe = get_safe(foods, possible, allergen_is)

    print(safe)
    tot = sum(ing in f for ing in safe for f in foods)
    print(f'total = {tot}')


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

