const std = @import("std");
const config = @import("config");

fn encode(in: std.ArrayList(u32), allocator: std.mem.Allocator) !std.ArrayList(u32) {
    var out = std.ArrayList(u32).init(allocator);

    var run_length: u32 = 1;
    for (1..in.items.len) |i| {
        const c = in.items[i];
        const p = in.items[i - 1];
        if (c == p) {
            run_length += 1;
        } else {
            try out.append(run_length);
            try out.append(p);
            run_length = 1;
        }
    }

    try out.append(run_length);
    try out.append(in.items[in.items.len - 1]);
    return out;
}

test "encode" {
    const allocator = std.testing.allocator;
    {
        const in = [_]u32{1};
        const out = [_]u32{ 1, 1 };

        var in_list = std.ArrayList(u32).init(allocator);
        defer in_list.deinit();

        for (0..in.len) |i| {
            try in_list.append(in[i]);
        }

        const result = try encode(in_list, allocator);
        defer result.deinit();

        try std.testing.expectEqual(out.len, result.items.len);

        for (0..out.len) |i| {
            try std.testing.expectEqual(out[i], result.items[i]);
        }
    }
    {
        const in = [_]u32{ 1, 1 };
        const out = [_]u32{ 2, 1 };

        var in_list = std.ArrayList(u32).init(allocator);
        defer in_list.deinit();

        for (0..in.len) |i| {
            try in_list.append(in[i]);
        }

        const result = try encode(in_list, allocator);
        defer result.deinit();

        try std.testing.expectEqual(out.len, result.items.len);

        for (0..out.len) |i| {
            try std.testing.expectEqual(out[i], result.items[i]);
        }
    }
    {
        const in = [_]u32{ 2, 1 };
        const out = [_]u32{ 1, 2, 1, 1 };

        var in_list = std.ArrayList(u32).init(allocator);
        defer in_list.deinit();

        for (0..in.len) |i| {
            try in_list.append(in[i]);
        }

        const result = try encode(in_list, allocator);
        defer result.deinit();

        try std.testing.expectEqual(out.len, result.items.len);

        for (0..out.len) |i| {
            try std.testing.expectEqual(out[i], result.items[i]);
        }
    }
    {
        const in = [_]u32{ 1, 2, 1, 1 };
        const out = [_]u32{ 1, 1, 1, 2, 2, 1 };

        var in_list = std.ArrayList(u32).init(allocator);
        defer in_list.deinit();

        for (0..in.len) |i| {
            try in_list.append(in[i]);
        }

        const result = try encode(in_list, allocator);
        defer result.deinit();

        try std.testing.expectEqual(out.len, result.items.len);

        for (0..out.len) |i| {
            try std.testing.expectEqual(out[i], result.items[i]);
        }
    }
    {
        const in = [_]u32{ 1, 1, 1, 2, 2, 1 };
        const out = [_]u32{ 3, 1, 2, 2, 1, 1 };

        var in_list = std.ArrayList(u32).init(allocator);
        defer in_list.deinit();

        for (0..in.len) |i| {
            try in_list.append(in[i]);
        }

        const result = try encode(in_list, allocator);
        defer result.deinit();

        try std.testing.expectEqual(out.len, result.items.len);

        for (0..out.len) |i| {
            try std.testing.expectEqual(out[i], result.items[i]);
        }
    }
}

pub fn a(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const s = std.mem.trim(u8, input_text, " \n");
    var digits = std.ArrayList(u32).init(allocator);

    for (0..s.len) |i| {
        try digits.append(s[i] - 48);
    }

    for (0..40) |_| {
        const out = try encode(digits, allocator);
        digits.deinit();
        digits = out;
    }
    defer digits.deinit();

    var len: u32 = 0;

    for (digits.items) |d| {
        len += std.math.log10(d) + 1;
    }

    if (!config.benchmark) {
        std.debug.print("day 10 a: {d}\n", .{len});
    }
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const s = std.mem.trim(u8, input_text, " \n");
    var digits = std.ArrayList(u32).init(allocator);

    for (0..s.len) |i| {
        try digits.append(s[i] - 48);
    }

    for (0..50) |_| {
        const out = try encode(digits, allocator);
        digits.deinit();
        digits = out;
    }
    defer digits.deinit();

    var len: u32 = 0;

    for (digits.items) |d| {
        len += std.math.log10(d) + 1;
    }

    if (!config.benchmark) {
        std.debug.print("day 10 b: {d}\n", .{len});
    }
}
