import aoc

import dataclasses
import enum
import logging

from typing import Any, List, Tuple, Dict, Optional


def part1(text: str) -> int:
    min_x = float("infinity")
    min_y = float("infinity")
    max_x = float("-infinity")
    max_y = float("-infinity")

    sensors = set()
    beacons = set()
    coverage = dict()

    for line in text.splitlines():
        toks = line.split()

        sensor_x = int(toks[2].split("=")[-1].strip(","))
        sensor_y = int(toks[3].split("=")[-1].strip(":"))
        sensor = sensor_x + (sensor_y * 1j)
        sensors.add(sensor)

        beacon_x = int(toks[8].split("=")[-1].strip(","))
        beacon_y = int(toks[9].split("=")[-1].strip(":"))
        beacon = beacon_x + (beacon_y * 1j)
        beacons.add(beacon)

        distance = beacon - sensor
        c = int(abs(distance.real) + abs(distance.imag))
        min_x = min(min_x, sensor_x - c)
        min_y = min(min_y, sensor_y - c)
        max_x = max(max_x, sensor_x + c)
        max_y = max(max_y, sensor_y + c)
        coverage[sensor] = c

    count = 0
    for j in range(min_x, max_x+1):
        p = j + (2000000 * 1j)
        if p in beacons or p in sensors:
            continue
        for sensor in sensors:
            p_dist = p - sensor
            p_dist = int(abs(p_dist.real) + abs(p_dist.imag))

            if p_dist <= coverage[sensor]:
                count += 1
                break

    return count

def part2(text: str) -> int:
    min_x = float("infinity")
    min_y = float("infinity")
    max_x = float("-infinity")
    max_y = float("-infinity")

    sensors = set()
    beacons = set()
    coverage = dict()

    for line in text.splitlines():
        toks = line.split()

        sensor_x = int(toks[2].split("=")[-1].strip(","))
        sensor_y = int(toks[3].split("=")[-1].strip(":"))
        sensor = sensor_x + (sensor_y * 1j)
        sensors.add(sensor)

        beacon_x = int(toks[8].split("=")[-1].strip(","))
        beacon_y = int(toks[9].split("=")[-1].strip(":"))
        beacon = beacon_x + (beacon_y * 1j)
        beacons.add(beacon)

        distance = beacon - sensor
        c = int(abs(distance.real) + abs(distance.imag))
        min_x = min(min_x, sensor_x - c)
        min_y = min(min_y, sensor_y - c)
        max_x = max(max_x, sensor_x + c)
        max_y = max(max_y, sensor_y + c)
        coverage[sensor] = c

    eligible = set()
    for sensor in sensors:
        c = coverage[sensor]+1
        for i in range(c+1):
            for p in [
                    (sensor + (c-i + (i * 1j))),
                    (sensor + (i-c + (i * 1j))),
                    (sensor + (c-i + (-i * 1j))),
                    (sensor + (i-c + (-i * 1j))),
            ]:
                if p.real < 0 or p.real > 4000000 or p.imag < 0 or p.imag > 4000000:
                    continue
                if p in sensors or p in beacons:
                    continue
                eligible.add(p)

    for p in eligible:
        found = True
        for sensor in sensors:
            p_dist = p - sensor
            p_dist = int(abs(p_dist.real) + abs(p_dist.imag))

            if p_dist <= coverage[sensor]:
                found = False
                break

        if found:
            return int(p.real * 4000000 + p.imag)

    return -1

def main():
    d = 15
    text = aoc.get_input(2022, d).rstrip()

    print(f'{d}.1: {part1(text)}')
    print(f'{d}.2: {part2(text)}')

    if logging.getLogger().level < logging.WARNING:
        tests: List[str] = [
        ]
        for test in tests:
            logging.info(f'test: "{test.rstrip()}"\n  {d}.1: {part1(test.rstrip())}\n  {d}.2: {part2(test.rstrip())}')

if __name__ == "__main__":
    main()

