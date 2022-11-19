const std = @import("std");
const utils = @import("./utils.zig");

const INPUT =
    \\input
;

pub fn part1(input: []const u8) isize {
    _ = input;
    return 0;
}

pub fn part2(input: []const u8) isize {
    _ = input;
    return 0;
}

pub fn main() void {
    std.debug.print("Part 1: {d}\n", .{part1(INPUT)});
    std.debug.print("Part 2: {d}\n", .{part2(INPUT)});
}

test "part1" {
    const cases: [0]utils.Case = .{};

    for (cases) |case| {
        try std.testing.expectEqual(case.expected, part1(case.input));
    }
}

test "part2" {
    const cases: [0]utils.Case = .{};

    for (cases) |case| {
        try std.testing.expectEqual(case.expected, part2(case.input));
    }
}
