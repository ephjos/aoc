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

    def get_data(self):
        return deepcopy(self._data)

    def __len__(self):
        return len(self._data)

    def __repr__(self):
        return repr(self._data)


class Deck(Queue):
    ''' A deck is just a Queue '''
    def __init__(self, data=[]):
        Queue.__init__(self, data)

    def copy(self):
        return Deck(deepcopy(self._data))

    def copy_next(self, i):
        return Deck(deepcopy(self._data[:i]))

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


P1_WIN = 1234
P2_WIN = 4321


def play_game(p1, p2, top=False):
    p1_mem = []
    p2_mem = []
    while not p1.empty() and not p2.empty():
        if top:
            print(f'p1 = {p1.score():10} p2 = {p2.score():10}')
        p1_data = p1.get_data()
        p2_data = p2.get_data()

        if p1_data in p1_mem or p2_data in p2_mem:
            return P1_WIN, p1

        p1_mem.append(p1_data)
        p2_mem.append(p2_data)

        a = p1.dequeue()
        b = p2.dequeue()
        u = len(p1)
        v = len(p2)
        curr_winner = None

        if a <= u and b <= v:
            curr_winner, _ = play_game(p1.copy_next(a), p2.copy_next(b))
        else:
            curr_winner = P1_WIN if a > b else P2_WIN

        if curr_winner == P1_WIN:
            p1.enqueue(a)
            p1.enqueue(b)
        else:
            p2.enqueue(b)
            p2.enqueue(a)

    if not p1.empty():
        return P1_WIN, p1
    else:
        return P2_WIN, p2


def main(fn):
    with open(fn, 'r') as fp:
        p1, p2 = build_decks(fp)
        win, deck = play_game(p1, p2, top=True)

        if win == P1_WIN:
            print('Winner: Player 1')
        else:
            print('Winner: Player 2')

        print(deck)
        print(f'score = {deck.score()}')


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file',
                        type=str, default='input', required=False)
    args = parser.parse_args()
    main(args.file)
