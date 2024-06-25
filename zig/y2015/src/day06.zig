const std = @import("std");
const config = @import("config");

pub fn a(input_text: []const u8, _: std.mem.Allocator) !void {
    var lights = std.mem.zeroes([1000][1000]u32);

    var in_num = false;
    var nums = std.mem.zeroes([4]u32);
    var ni: usize = 0;

    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        in_num = false;
        nums = std.mem.zeroes([4]u32);
        ni = 0;
        for (line) |c| {
            if (std.ascii.isDigit(c)) {
                nums[ni] *= 10;
                nums[ni] += (c - 48);
                in_num = true;
            } else if (in_num) {
                ni += 1;
                in_num = false;
            }
        }

        switch (line[6]) {
            'n' => {
                // On
                for (nums[1]..nums[3] + 1) |i| {
                    for (nums[0]..nums[2] + 1) |j| {
                        lights[i][j] = 1;
                    }
                }
            },
            'f' => {
                // Off
                for (nums[1]..nums[3] + 1) |i| {
                    for (nums[0]..nums[2] + 1) |j| {
                        lights[i][j] = 0;
                    }
                }
            },
            ' ' => {
                // Toggle
                for (nums[1]..nums[3] + 1) |i| {
                    for (nums[0]..nums[2] + 1) |j| {
                        lights[i][j] = 1 - lights[i][j];
                    }
                }
            },
            else => @panic("Unknown char"),
        }
    }

    var sum: u32 = 0;
    for (0..1000) |i| {
        for (0..1000) |j| {
            sum += lights[i][j];
        }
    }

    if (!config.benchmark) {
        std.debug.print("day  6 a: {d}\n", .{sum});
    }
}

pub fn b(input_text: []const u8, _: std.mem.Allocator) !void {
    var lights = std.mem.zeroes([1000][1000]u32);

    var in_num = false;
    var nums = std.mem.zeroes([4]u32);
    var ni: usize = 0;

    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        in_num = false;
        nums = std.mem.zeroes([4]u32);
        ni = 0;
        for (line) |c| {
            if (std.ascii.isDigit(c)) {
                nums[ni] *= 10;
                nums[ni] += (c - 48);
                in_num = true;
            } else if (in_num) {
                ni += 1;
                in_num = false;
            }
        }

        switch (line[6]) {
            'n' => {
                // On
                for (nums[1]..nums[3] + 1) |i| {
                    for (nums[0]..nums[2] + 1) |j| {
                        lights[i][j] += 1;
                    }
                }
            },
            'f' => {
                // Off
                for (nums[1]..nums[3] + 1) |i| {
                    for (nums[0]..nums[2] + 1) |j| {
                        if (lights[i][j] == 0) continue;
                        lights[i][j] -= 1;
                    }
                }
            },
            ' ' => {
                // Toggle
                for (nums[1]..nums[3] + 1) |i| {
                    for (nums[0]..nums[2] + 1) |j| {
                        lights[i][j] += 2;
                    }
                }
            },
            else => @panic("Unknown char"),
        }
    }

    var sum: u32 = 0;
    for (0..1000) |i| {
        for (0..1000) |j| {
            sum += lights[i][j];
        }
    }

    if (!config.benchmark) {
        std.debug.print("day  6 b: {d}\n", .{sum});
    }
}
