const std = @import("std");
const config = @import("config");

fn a__is_nice(s: []const u8) bool {
    var vowel_count: u32 = 0;
    var double_count: u32 = 0;
    var bad_double_count: u32 = 0;

    var prev: u8 = 0;

    for (s) |c| {
        switch (c) {
            'a', 'e', 'i', 'o', 'u' => vowel_count += 1,
            'b' => {
                bad_double_count += @intFromBool(prev == 'a');
            },
            'd' => {
                bad_double_count += @intFromBool(prev == 'c');
            },
            'q' => {
                bad_double_count += @intFromBool(prev == 'p');
            },
            'y' => {
                bad_double_count += @intFromBool(prev == 'x');
            },
            else => {},
        }

        if (c == prev) double_count += 1;
        prev = c;
    }

    return vowel_count >= 3 and double_count >= 1 and bad_double_count == 0;
}

test "a__is_nice" {
    try std.testing.expectEqual(true, a__is_nice("ugknbfddgicrmopn"));
    try std.testing.expectEqual(true, a__is_nice("aaa"));

    try std.testing.expectEqual(false, a__is_nice("jchzalrnumimnmhp"));
    try std.testing.expectEqual(false, a__is_nice("haegwjzuvuyypxyu"));
    try std.testing.expectEqual(false, a__is_nice("dvszwmarrgswjxmb"));
}

pub fn a(input_text: []const u8, _: std.mem.Allocator) !void {
    var nice_count: u32 = 0;
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        nice_count += @intFromBool(a__is_nice(line));
    }

    if (!config.benchmark) {
        std.debug.print("day  5 a: {d}\n", .{nice_count});
    }
}

fn b__is_nice(s: []const u8) bool {
    // Input strings are 16 characters long, there can only be 15 ordered pairs
    const MAX_PAIRS = 15;

    // Pack two chars (u8) into a u16, first u8 gets high bits
    var pairs: [MAX_PAIRS]u16 = undefined;
    var pairs_pos: usize = 0;

    var pair_no_overlap_count: u32 = 0;
    var xyx_count: u32 = 0;

    var prev: u8 = 0;
    var prev_prev: u8 = 0;

    for (s, 0..) |c, i| {
        const curr_pair = (@as(u16, prev) << 8) | c;
        for (0..pairs_pos) |j| {
            pair_no_overlap_count += @intFromBool(pairs[j] == curr_pair and i - j > 2);
        }

        xyx_count += @intFromBool(c == prev_prev);

        if (i >= 1) {
            pairs[pairs_pos] = (@as(u16, s[i - 1]) << 8) | (s[i]);
            pairs_pos += 1;
        }

        prev_prev = prev;
        prev = c;
    }

    return pair_no_overlap_count >= 1 and xyx_count >= 1;
}

test "b__is_nice" {
    try std.testing.expectEqual(true, b__is_nice("qjhvhtzxzqqjkmpb"));
    try std.testing.expectEqual(true, b__is_nice("xxyxx"));

    try std.testing.expectEqual(false, b__is_nice("uurcxstgmygtbstg"));
    try std.testing.expectEqual(false, b__is_nice("ieodomkazucvgmuy"));
}

pub fn b(input_text: []const u8, _: std.mem.Allocator) !void {
    var nice_count: u32 = 0;
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        nice_count += @intFromBool(b__is_nice(line));
    }

    if (!config.benchmark) {
        std.debug.print("day  5 b: {d}\n", .{nice_count});
    }
}
