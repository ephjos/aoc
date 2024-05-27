const std = @import("std");
const config = @import("config");
const input_text = @embedFile("./inputs/d01");

pub fn a(_: std.mem.Allocator) void {
    var floor: i32 = 0;
    for (input_text) |c| {
        switch (c) {
            '(' => floor += 1,
            ')' => floor -= 1,
            else => {},
        }
    }

    if (!config.benchmark) {
        std.debug.print("day  1 a: {d}\n", .{floor});
    }
}

pub fn b(_: std.mem.Allocator) void {
    var floor: i32 = 0;
    var position: usize = 0;
    for (input_text, 0..) |c, i| {
        switch (c) {
            '(' => floor += 1,
            ')' => floor -= 1,
            else => {},
        }
        if (floor < 0) {
            position = i + 1;
            break;
        }
    }

    if (!config.benchmark) {
        std.debug.print("day  1 b: {d}\n", .{position});
    }
}
