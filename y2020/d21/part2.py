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

    safe_set = set(safe)
    food_set = set()
    for f in foods:
        food_set |= f
    dangerous_ings = list(food_set - safe_set)
    n = len(dangerous_ings)
    dangerous_list = []
    while len(dangerous_list) != n:
        for i, ding in enumerate(dangerous_ings):
            alls = possible[ding]
            if len(alls) == 1:
                v = list(alls)[0]
                dangerous_list.append((ding, v))
                for ing in dangerous_ings:
                    possible[ing] -= set([v])
                del dangerous_ings[i]
                break

    print(','.join([x[0] for x in sorted(dangerous_list, key=lambda x: x[1])]))

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)

