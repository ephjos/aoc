const std = @import("std");
const config = @import("config");

fn step(in: *[8]u8) !void {
    inline for (0..8) |i| {
        const j = 8 - i - 1;
        in[j] = (((in[j] + 1) - 97) % 26) + 97;
        if (in[j] != 97) {
            break;
        }
    }

    return;
}

fn check(in: [8]u8) !bool {
    var has_straight_3 = false;

    inline for (0..5) |i| {
        has_straight_3 = has_straight_3 or (in[i] + 1 == in[i + 1]) and (in[i + 1] + 1 == in[i + 2]);
    }

    var has_iol = false;
    inline for (0..8) |i| {
        has_iol = has_iol or in[i] == 'i' or in[i] == 'o' or in[i] == 'l';
    }

    var has_pairs: u32 = 0;
    var pi: usize = 0;
    while (pi < 7) {
        if (in[pi] == in[pi + 1]) {
            has_pairs += 1;
            pi += 2;
        } else {
            pi += 1;
        }
    }

    return has_straight_3 and !has_iol and (has_pairs >= 2);
}

pub fn a(input_text: []const u8, _: std.mem.Allocator) !void {
    const s = std.mem.trim(u8, input_text, " \n");
    var digits = [_]u8{0} ** 8;

    for (0..s.len) |i| {
        digits[i] = s[i];
    }

    while (!try check(digits)) {
        try step(&digits);
    }

    if (!config.benchmark) {
        std.debug.print("day 11 a: {s}\n", .{digits});
    }
}

pub fn b(input_text: []const u8, _: std.mem.Allocator) !void {
    const s = std.mem.trim(u8, input_text, " \n");
    var digits = [_]u8{0} ** 8;

    for (0..s.len) |i| {
        digits[i] = s[i];
    }

    while (!try check(digits)) {
        try step(&digits);
    }

    try step(&digits);
    while (!try check(digits)) {
        try step(&digits);
    }

    if (!config.benchmark) {
        std.debug.print("day 11 b: {s}\n", .{digits});
    }
}
