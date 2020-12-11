#!/usr/bin/env python

def trib(n):
    if n == 0:
        return 0
    if n == 1 or n == 2:
        return 1

    return trib(n-1)+trib(n-2)+trib(n-3)

def main():
    with open("./input", "r") as fp:
        nums = [0]
        for line in fp:
            nums.append(int(line))

        nums = sorted(nums)
        nums.append(nums[-1]+3)
        print(nums)

        diffs = [nums[i+1] - nums[i] for i in range(len(nums)-1)]
        print(diffs)

        r_diffs = ''.join(map(lambda x: str(x), diffs)).split('3')
        rr_diffs = [len(x)+1 for x in r_diffs if x]
        print(rr_diffs)

        r = 1
        for rr in [trib(x) for x in rr_diffs]:
            r *= rr

        print(r)

if __name__ == "__main__":
    main()

