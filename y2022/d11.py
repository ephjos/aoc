import aoc

import dataclasses
import enum
import logging

from typing import Any, Callable, List, Tuple, Dict, Optional
@dataclasses.dataclass
class Monkey:
    items: List[int]
    #operation: Callable[[int], int]
    operation_toks: List[str]
    #test: Callable[[int], bool]
    test_d: int
    true_target: int
    false_target: int

    def operation(self, x: int) -> int:
        if self.operation_toks[1] == "+":
            if self.operation_toks[2] == "old":
                return x + x
            else:
                return x + int(self.operation_toks[2])
        else:
            if self.operation_toks[2] == "old":
                return x * x
            else:
                return x * int(self.operation_toks[2])

    def test(self, x: int) -> bool:
        return x % self.test_d == 0

def parse_monkeys(text: str) -> List[Monkey]:
    monkeys = []
    for block in text.split("\n\n"):
        name, items_line, operation_line, test_line, true_line, false_line = block.splitlines()

        items = [int(item_str) for item_str in items_line.split(":")[1].strip().split(", ")]

        operation_toks = operation_line.split("=")[1].strip().split()

        test_d = int(test_line.split()[-1])

        true_target = int(true_line.split()[-1])
        false_target = int(false_line.split()[-1])

        monkeys.append(Monkey(
            items=items,
            operation_toks=operation_toks,
            test_d=test_d,
            true_target=true_target,
            false_target=false_target,
        ))
    return monkeys

def part1(text: str) -> int:
    monkeys = parse_monkeys(text)
    rounds = 20
    inspections_by_monkey = [0 for _ in range(len(monkeys))]

    for r in range(rounds):
        for i in range(len(monkeys)):
            monkey = monkeys[i]
            while monkey.items:
                item = monkey.items.pop(0)
                worry = monkey.operation(item)
                inspections_by_monkey[i] += 1
                worry = worry // 3
                if monkey.test(worry):
                    monkeys[monkey.true_target].items.append(worry)
                else:
                    monkeys[monkey.false_target].items.append(worry)

    sorted_counts = list(sorted(inspections_by_monkey, reverse=True))[0:2]
    return sorted_counts[0] * sorted_counts[1]

def part2(text: str) -> int:
    monkeys = parse_monkeys(text)
    rounds = 10000
    inspections_by_monkey = [0 for _ in range(len(monkeys))]

    lcm = 1
    for i in range(len(monkeys)):
        lcm *= monkeys[i].test_d

    for r in range(rounds):
        for i in range(len(monkeys)):
            monkey = monkeys[i]
            while monkey.items:
                item = monkey.items.pop(0)
                worry = monkey.operation(item)
                inspections_by_monkey[i] += 1
                if monkey.test(worry):
                    target = monkeys[monkey.true_target]
                else:
                    target = monkeys[monkey.false_target]
                target.items.append(worry % lcm)

    sorted_counts = list(sorted(inspections_by_monkey, reverse=True))[0:2]
    return sorted_counts[0] * sorted_counts[1]

def main():
    d = 11
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
"""\
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
"""
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

