const std = @import("std");
const config = @import("config");

pub fn a(input_text: []const u8, _: std.mem.Allocator) !void {
    var sum: i32 = 0;
    var num: i32 = 0;
    var sign: i2 = 1;

    for (input_text) |c| {
        switch (c) {
            '-' => {
                sign = -1;
            },
            '0'...'9' => {
                num *= 10;
                num += c - 48;
            },
            else => {
                sum += sign * num;
                num = 0;
                sign = 1;
            },
        }
    }

    if (!config.benchmark) {
        std.debug.print("day 12 a: {d}\n", .{sum});
    }
}

fn b__count(json: std.json.Value) !i32 {
    var sum: i32 = 0;
    switch (json) {
        .null, .bool, .number_string, .string => {},
        .float => {
            sum += @as(i32, @intFromFloat(json.float));
        },
        .integer => {
            sum += @intCast(json.integer);
        },
        .array => {
            for (json.array.items) |v| {
                sum += try b__count(v);
            }
        },
        .object => {
            var in_red = false;
            var e_it = json.object.iterator();
            while (e_it.next()) |e| {
                const v = e.value_ptr.*;
                switch (v) {
                    .string => {
                        if (std.mem.eql(u8, "red", v.string)) {
                            in_red = true;
                        }
                    },
                    else => {},
                }
            }

            if (!in_red) {
                e_it = json.object.iterator();
                while (e_it.next()) |e| {
                    sum += try b__count(e.value_ptr.*);
                }
            }
        },
    }

    return sum;
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const parsed = try std.json.parseFromSlice(std.json.Value, allocator, input_text, .{});
    defer parsed.deinit();

    const sum = try b__count(parsed.value);

    if (!config.benchmark) {
        std.debug.print("day 12 b: {d}\n", .{sum});
    }
}
