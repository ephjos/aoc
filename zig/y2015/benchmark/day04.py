import subprocess
import sys

for day04_block_size in range(8, 32, 2):
    cmd = f"""zig build -Dbenchmark=true -Doptimize=ReleaseFast -Dday04_block_size={1<<day04_block_size} && ./zig-out/bin/y2015 --day 4 2>&1 | awk -F" " '{{print $4}}'"""
    output = subprocess.check_output(cmd, shell=True).decode().strip()[:-2]
    print(f"{day04_block_size} {output}", flush=True)
