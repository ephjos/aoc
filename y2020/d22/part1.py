#!/usr/bin/env python
import argparse
from copy import deepcopy


class Queue:
    def __init__(self, data=[]):
        self._data = deepcopy(data)

    def enqueue(self, x):
        self._data.append(x)

    def top(self):
        if self._data != []:
            return self._data[0]
        return None

    def dequeue(self):
        if self._data != []:
            return self._data.pop(0)
        return None

    def empty(self):
        return self._data == []

    def __len__(self):
        return len(self._data)

    def __repr__(self):
        return repr(self._data)


class Deck(Queue):
    ''' A deck is just a Queue '''
    def __init__(self, data=[]):
        Queue.__init__(self, data)

    def score(self):
        s = 0
        temp_data = deepcopy(self._data)
        while not self.empty():
            s += len(self) * self.top()
            self.dequeue()
        self._data = temp_data
        return s


def build_deck(lines):
    cards = [int(card) for card in lines[1:]]
    return Deck(cards)


def build_decks(fp):
    p1_lines, p2_lines = (block.strip().split('\n')
                          for block in fp.read().split('\n\n'))
    return build_deck(p1_lines), build_deck(p2_lines)


def play_game(p1, p2):
    while not p1.empty() and not p2.empty():
        a = p1.dequeue()
        b = p2.dequeue()
        if a > b:
            p1.enqueue(a)
            p1.enqueue(b)
        else:
            p2.enqueue(b)
            p2.enqueue(a)

    if not p1.empty():
        print('Winner: Player 1')
        print(p1)
        print(f'Score = {p1.score()}')
    else:
        print('Winner: Player 2')
        print(p2)
        print(f'Score = {p2.score()}')


def main(fn):
    with open(fn, 'r') as fp:
        p1, p2 = build_decks(fp)
        play_game(p1, p2)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)
