const std = @import("std");
const config = @import("config");

const Day01 = @import("./day01.zig");
const Day02 = @import("./day02.zig");
const Day03 = @import("./day03.zig");
const Day04 = @import("./day04.zig");
const Day05 = @import("./day05.zig");
const Day06 = @import("./day06.zig");
const Day07 = @import("./day07.zig");
const Day08 = @import("./day08.zig");
const Day09 = @import("./day09.zig");
const Day10 = @import("./day10.zig");
const Day11 = @import("./day11.zig");
const Day12 = @import("./day12.zig");
const Day13 = @import("./day13.zig");

const NANOS_IN_MILI: f64 = 1_000_000;

const ExitCode = enum(u8) { OK, NO_VALUE_FOR_ARGUMENT };
const DayPart = enum(u8) { A, B };

test "check mem" {
    try run(std.testing.allocator);
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

fn run_day_part(d: u8, p: DayPart, inputs: std.ArrayList([]u8), allocator: std.mem.Allocator) anyerror!void {
    switch (d) {
        inline 1 => if (p == DayPart.A) try Day01.a(inputs.items[d - 1], allocator) else try Day01.b(inputs.items[d - 1], allocator),
        inline 2 => if (p == DayPart.A) try Day02.a(inputs.items[d - 1], allocator) else try Day02.b(inputs.items[d - 1], allocator),
        inline 3 => if (p == DayPart.A) try Day03.a(inputs.items[d - 1], allocator) else try Day03.b(inputs.items[d - 1], allocator),
        inline 4 => if (p == DayPart.A) try Day04.a(inputs.items[d - 1], allocator) else try Day04.b(inputs.items[d - 1], allocator),
        inline 5 => if (p == DayPart.A) try Day05.a(inputs.items[d - 1], allocator) else try Day05.b(inputs.items[d - 1], allocator),
        inline 6 => if (p == DayPart.A) try Day06.a(inputs.items[d - 1], allocator) else try Day06.b(inputs.items[d - 1], allocator),
        inline 7 => if (p == DayPart.A) try Day07.a(inputs.items[d - 1], allocator) else try Day07.b(inputs.items[d - 1], allocator),
        inline 8 => if (p == DayPart.A) try Day08.a(inputs.items[d - 1], allocator) else try Day08.b(inputs.items[d - 1], allocator),
        inline 9 => if (p == DayPart.A) try Day09.a(inputs.items[d - 1], allocator) else try Day09.b(inputs.items[d - 1], allocator),
        inline 10 => if (p == DayPart.A) try Day10.a(inputs.items[d - 1], allocator) else try Day10.b(inputs.items[d - 1], allocator),
        inline 11 => if (p == DayPart.A) try Day11.a(inputs.items[d - 1], allocator) else try Day11.b(inputs.items[d - 1], allocator),
        inline 12 => if (p == DayPart.A) try Day12.a(inputs.items[d - 1], allocator) else try Day12.b(inputs.items[d - 1], allocator),
        inline 13 => if (p == DayPart.A) try Day13.a(inputs.items[d - 1], allocator) else try Day13.b(inputs.items[d - 1], allocator),
        inline else => {},
    }
}

fn run_day(d: u8, inputs: std.ArrayList([]u8), allocator: std.mem.Allocator) anyerror!f64 {
    // No benchmarking, just run
    if (!config.benchmark) {
        try run_day_part(d, DayPart.A, inputs, allocator);
        try run_day_part(d, DayPart.B, inputs, allocator);
        return 0;
    } else {
        var timer = try std.time.Timer.start();
        const samples_f64 = @as(f64, @floatFromInt(config.samples));

        // Time part a
        const a_start = timer.read();
        @setEvalBranchQuota(1 << 14);
        inline for (0..config.samples) |_| {
            try run_day_part(d, DayPart.A, inputs, allocator);
        }
        const a_end = timer.read();
        const a_time: f64 = (@as(f64, @floatFromInt(a_end - a_start)) / samples_f64) / NANOS_IN_MILI;

        // Time part b
        const b_start = timer.read();
        @setEvalBranchQuota(1 << 14);
        inline for (0..config.samples) |_| {
            try run_day_part(d, DayPart.B, inputs, allocator);
        }
        const b_end = timer.read();
        const b_time: f64 = (@as(f64, @floatFromInt(b_end - b_start)) / samples_f64) / NANOS_IN_MILI;

        // Report
        const total = a_time + b_time;
        std.debug.print("day {d:2} {d:10.2}ms {d:10.2}ms {d:10.2}ms\n", .{ d, a_time, b_time, total });
        return total;
    }
}

fn get_inputs(allocator: std.mem.Allocator) !std.ArrayList([]u8) {
    const MAX_FILE_BYTES = 1 << 32;
    var name_buf: [64]u8 = undefined;
    var list = std.ArrayList([]u8).init(allocator);
    const cwd = std.fs.cwd();
    const inputs_dir = try cwd.openDir("inputs", .{});
    for (1..26) |i| {
        if (i < 10) {
            const path = try std.fmt.bufPrint(&name_buf, "d0{d}", .{i});
            try list.append(try inputs_dir.readFileAlloc(allocator, path, MAX_FILE_BYTES));
        } else {
            const path = try std.fmt.bufPrint(&name_buf, "d{d}", .{i});
            try list.append(try inputs_dir.readFileAlloc(allocator, path, MAX_FILE_BYTES));
        }
    }

    return list;
}

pub fn run(allocator: std.mem.Allocator) !void {
    var args = try std.process.ArgIterator.initWithAllocator(allocator);

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

    const inputs = try get_inputs(allocator);
    defer {
        for (inputs.items) |item| {
            allocator.free(item);
        }
        inputs.deinit();
    }

    // Run single day
    if (arguments.day != null) {
        _ = try run_day(arguments.day.?, inputs, allocator);
        return;
    }

    // Run all
    var sum: f64 = 0;
    inline for (1..25) |i| {
        sum += run_day(@as(u8, @intCast(i)), inputs, allocator) catch 0;
    }

    if (config.benchmark) {
        std.debug.print("total: {d:10.2}ms\n", .{sum});
    }
}

pub fn main() !void {
    try run(std.heap.page_allocator);
}
