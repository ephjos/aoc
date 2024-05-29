import subprocess
import sys

for day04_window_size in range(28, 12, -2):
    for day04_block_size in range(day04_window_size-2, 8, -2):
        cmd = f"""zig build -Dbenchmark=true -Doptimize=ReleaseFast -Dday04_window_size={1<<day04_window_size} -Dday04_block_size={1<<day04_block_size} && ./zig-out/bin/y2015 --day 4 2>&1 | awk -F" " '{{print $4}}'"""
        #print(f"Running: {cmd}", file=sys.stderr)
        output = subprocess.check_output(cmd, shell=True).decode().strip()[:-2]
        print(f"{1<<day04_window_size}_{1<<day04_block_size} {output}", flush=True)
