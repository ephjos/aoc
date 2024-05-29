const std = @import("std");
const config = @import("config");

const Day01 = @import("./day01.zig");
const Day02 = @import("./day02.zig");
const Day03 = @import("./day03.zig");
const Day04 = @import("./day04.zig");

const NANOS_IN_MILI: f64 = 1_000_000;

const ExitCode = enum(u8) { OK, NO_VALUE_FOR_ARGUMENT };
const DayPart = enum(u8) { A, B };

test "check mem" {
    const allocator = std.testing.allocator;

    inline for (1..25) |i| {
        _ = run_day(@as(u8, @intCast(i)), allocator) catch 0;
    }
}

const Arguments = struct {
    day: ?u8,

    pub fn isHelp(arg: []const u8) bool {
        return std.mem.eql(u8, arg, "-h") or std.mem.eql(u8, arg, "--help");
    }

    pub fn isDay(arg: []const u8) bool {
        return std.mem.eql(u8, arg, "-d") or std.mem.eql(u8, arg, "--day");
    }

    pub fn getValueFor(arg: []const u8, val: ?[]const u8) ?[]const u8 {
        if (val == null) {
            std.debug.print("No value provided for {s}\n\n{s}\n", .{ arg, HELP_TEXT });
            std.process.exit(@intFromEnum(ExitCode.NO_VALUE_FOR_ARGUMENT));
        }

        return val;
    }
};

const HELP_TEXT = @embedFile("./help.txt");

fn run_day_part(d: u8, p: DayPart, allocator: std.mem.Allocator) anyerror!void {
    switch (d) {
        inline 1 => if (p == DayPart.A) Day01.a(allocator) else Day01.b(allocator),
        inline 2 => if (p == DayPart.A) try Day02.a(allocator) else try Day02.b(allocator),
        inline 3 => if (p == DayPart.A) try Day03.a(allocator) else try Day03.b(allocator),
        inline 4 => if (p == DayPart.A) try Day04.a(allocator) else try Day04.b(allocator),
        inline else => {},
    }
}

fn run_day(d: u8, allocator: std.mem.Allocator) anyerror!f64 {
    // No benchmarking, just run
    if (!config.benchmark) {
        try run_day_part(d, DayPart.A, allocator);
        try run_day_part(d, DayPart.B, allocator);
        return 0;
    } else {
        var timer = try std.time.Timer.start();
        const samples_f64 = @as(f64, @floatFromInt(config.samples));

        // Time part a
        const a_start = timer.read();
        @setEvalBranchQuota(1 << 14);
        inline for (0..config.samples) |_| {
            try run_day_part(d, DayPart.A, allocator);
        }
        const a_end = timer.read();
        const a_time: f64 = (@as(f64, @floatFromInt(a_end - a_start)) / samples_f64) / NANOS_IN_MILI;

        // Time part b
        const b_start = timer.read();
        @setEvalBranchQuota(1 << 14);
        inline for (0..config.samples) |_| {
            try run_day_part(d, DayPart.B, allocator);
        }
        const b_end = timer.read();
        const b_time: f64 = (@as(f64, @floatFromInt(b_end - b_start)) / samples_f64) / NANOS_IN_MILI;

        // Report
        const total = a_time + b_time;
        std.debug.print("day {d:2} {d:7.2}ms {d:7.2}ms {d:7.2}ms\n", .{ d, a_time, b_time, total });
        return total;
    }
}

pub fn main() !void {
    std.debug.print("\n", .{});
    const allocator = std.heap.page_allocator;

    var args = try std.process.ArgIterator.initWithAllocator(allocator);
    defer args.deinit();

    var arguments = Arguments{
        .day = null,
    };

    while (args.next()) |arg| {
        if (Arguments.isHelp(arg)) {
            std.debug.print("{s}\n", .{HELP_TEXT});
            std.process.exit(@intFromEnum(ExitCode.OK));
        } else if (Arguments.isDay(arg)) {
            const value = Arguments.getValueFor(arg, args.next());
            arguments.day = try std.fmt.parseInt(u8, value.?, 10);
        }
    }

    // Run single day
    if (arguments.day != null) {
        _ = try run_day(arguments.day.?, allocator);
        return;
    }

    // Run all
    var sum: f64 = 0;
    inline for (1..25) |i| {
        sum += run_day(@as(u8, @intCast(i)), allocator) catch 0;
    }

    if (config.benchmark) {
        std.debug.print("total: {d:7.2}ms\n", .{sum});
    }
}
