const std = @import("std");
const config = @import("config");

test "encode" {
    const allocator = std.testing.allocator;
    {
        const in = [_]u8{1};
        const expected = [_]u8{ 1, 1 };

        var out = try allocator.alloc(u8, 1 << 12);
        defer allocator.free(out);

        const out_len = try encode__(in.len, &in, &out, allocator);

        try std.testing.expectEqual(expected.len, out_len);

        for (0..out_len) |i| {
            try std.testing.expectEqual(expected[i], out[i]);
        }
    }
    {
        const in = [_]u8{ 1, 1 };
        const expected = [_]u8{ 2, 1 };

        var out = try allocator.alloc(u8, 1 << 12);
        defer allocator.free(out);

        const out_len = try encode__(in.len, &in, &out, allocator);

        try std.testing.expectEqual(expected.len, out_len);

        for (0..out_len) |i| {
            try std.testing.expectEqual(expected[i], out[i]);
        }
    }
    {
        const in = [_]u8{ 2, 1 };
        const expected = [_]u8{ 1, 2, 1, 1 };

        var out = try allocator.alloc(u8, 1 << 12);
        defer allocator.free(out);

        const out_len = try encode__(in.len, &in, &out, allocator);

        try std.testing.expectEqual(expected.len, out_len);

        for (0..out_len) |i| {
            try std.testing.expectEqual(expected[i], out[i]);
        }
    }
    {
        const in = [_]u8{ 1, 2, 1, 1 };
        const expected = [_]u8{ 1, 1, 1, 2, 2, 1 };

        var out = try allocator.alloc(u8, 1 << 12);
        defer allocator.free(out);

        const out_len = try encode__(in.len, &in, &out, allocator);

        try std.testing.expectEqual(expected.len, out_len);

        for (0..out_len) |i| {
            try std.testing.expectEqual(expected[i], out[i]);
        }
    }
    {
        const in = [_]u8{ 1, 1, 1, 2, 2, 1 };
        const expected = [_]u8{ 3, 1, 2, 2, 1, 1 };

        var out = try allocator.alloc(u8, 1 << 12);
        defer allocator.free(out);

        const out_len = try encode__(in.len, &in, &out, allocator);

        try std.testing.expectEqual(expected.len, out_len);

        for (0..out_len) |i| {
            try std.testing.expectEqual(expected[i], out[i]);
        }
    }
}

fn encode__(in_len: usize, in: []const u8, out: *[]u8, _: std.mem.Allocator) !usize {
    var out_i: usize = 0;
    var run_length: u8 = 1;
    for (1..in_len) |i| {
        const c = in[i];
        const p = in[i - 1];
        if (c != p) {
            out.*[out_i] = run_length;
            out_i += 1;
            out.*[out_i] = p;
            out_i += 1;
            run_length = 0;
        }
        run_length += 1;
    }

    out.*[out_i] = run_length;
    out_i += 1;
    out.*[out_i] = in[in_len - 1];
    out_i += 1;
    return out_i;
}

pub fn a(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const s = std.mem.trim(u8, input_text, " \n");

    var x = try allocator.alloc(u8, 1 << 24);
    defer allocator.free(x);
    var x_len: usize = s.len;

    var y = try allocator.alloc(u8, 1 << 24);
    defer allocator.free(y);
    var y_len: usize = 0;

    for (0..x_len) |i| {
        x[i] = s[i] - 48;
    }

    inline for (0..20) |_| {
        y_len = try encode__(x_len, x, &y, allocator);
        x_len = try encode__(y_len, y, &x, allocator);
    }

    if (!config.benchmark) {
        std.debug.print("day 10 a: {d}\n", .{x_len});
    }
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const s = std.mem.trim(u8, input_text, " \n");

    var x = try allocator.alloc(u8, 1 << 24);
    defer allocator.free(x);
    var x_len: usize = s.len;

    var y = try allocator.alloc(u8, 1 << 24);
    defer allocator.free(y);
    var y_len: usize = 0;

    for (0..x_len) |i| {
        x[i] = s[i] - 48;
    }

    inline for (0..25) |_| {
        y_len = try encode__(x_len, x, &y, allocator);
        x_len = try encode__(y_len, y, &x, allocator);
    }

    if (!config.benchmark) {
        std.debug.print("day 10 b: {d}\n", .{x_len});
    }
}
