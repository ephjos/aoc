const std = @import("std");
const config = @import("config");

const StringData = struct {
    code_len: u32,
    mem_len: u32,
};

fn a__string_data(s: []const u8) !StringData {
    var count: u32 = 0;
    var escaping = false;
    var i: usize = 0;

    while (i < s.len) {
        const c = s[i];
        if (!escaping and (c == '\\')) {
            escaping = true;
            i += 1;
        } else if (escaping and (c == '"' or c == '\\')) {
            count += 1;
            escaping = false;
            i += 1;
        } else if (escaping and (c == 'x')) {
            count += 1;
            escaping = false;
            i += 3;
        } else if (!escaping and (c == '"')) {
            i += 1;
        } else {
            count += 1;
            i += 1;
        }
    }

    return .{
        .code_len = @intCast(s.len),
        .mem_len = count,
    };
}

test "a__string_data" {
    try std.testing.expectEqual(StringData{ .code_len = 2, .mem_len = 0 }, try a__string_data("\"\""));
    try std.testing.expectEqual(StringData{ .code_len = 5, .mem_len = 3 }, try a__string_data("\"abc\""));
    try std.testing.expectEqual(StringData{ .code_len = 10, .mem_len = 7 }, try a__string_data("\"aaa\\\"aaa\""));
    try std.testing.expectEqual(StringData{ .code_len = 6, .mem_len = 1 }, try a__string_data("\"\\x27\""));
}

pub fn a(input_text: []const u8, _: std.mem.Allocator) !void {
    var sum: u32 = 0;
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        const string_data = try a__string_data(line);
        sum += string_data.code_len - string_data.mem_len;
    }

    if (!config.benchmark) {
        std.debug.print("day  8 a: {d}\n", .{sum});
    }
}

fn b__string_data(s: []const u8) !StringData {
    var count: u32 = 2;

    for (s) |c| {
        switch (c) {
            '"', '\\' => count += 2,
            else => count += 1,
        }
    }

    return .{
        .code_len = @intCast(s.len),
        .mem_len = count,
    };
}

test "b__string_data" {
    try std.testing.expectEqual(StringData{ .code_len = 2, .mem_len = 6 }, try b__string_data("\"\""));
    try std.testing.expectEqual(StringData{ .code_len = 5, .mem_len = 9 }, try b__string_data("\"abc\""));
    try std.testing.expectEqual(StringData{ .code_len = 10, .mem_len = 16 }, try b__string_data("\"aaa\\\"aaa\""));
    try std.testing.expectEqual(StringData{ .code_len = 6, .mem_len = 11 }, try b__string_data("\"\\x27\""));
}

pub fn b(input_text: []const u8, _: std.mem.Allocator) !void {
    var sum: u32 = 0;
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        const string_data = try b__string_data(line);
        sum += string_data.mem_len - string_data.code_len;
    }

    if (!config.benchmark) {
        std.debug.print("day  8 b: {d}\n", .{sum});
    }
}
