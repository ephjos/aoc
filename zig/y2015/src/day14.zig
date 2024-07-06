const std = @import("std");
const config = @import("config");

fn a__core(input_text: []const u8, time_steps: u32, _: std.mem.Allocator) !u32 {
    const NUM_REINDEER = 9;

    var speeds = [_]u32{0} ** NUM_REINDEER;
    var durations = [_]u32{0} ** NUM_REINDEER;
    var rests = [_]u32{0} ** NUM_REINDEER;
    var i: usize = 0;

    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.split(u8, line, " ");
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        const speed_raw = toks.next().?;
        const speed = try std.fmt.parseInt(u8, speed_raw, 10);
        _ = toks.next().?;
        _ = toks.next().?;
        const duration_raw = toks.next().?;
        const duration = try std.fmt.parseInt(u8, duration_raw, 10);
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        const rest_raw = toks.next().?;
        const rest = try std.fmt.parseInt(u8, rest_raw, 10);

        speeds[i] = speed;
        durations[i] = duration;
        rests[i] = rest;
        i += 1;
    }

    var dists = [_]u32{0} ** NUM_REINDEER;

    for (0..i) |j| {
        const cycle_time = durations[j] + rests[j];
        const cycle_dist = speeds[j];
        dists[j] = (time_steps / cycle_time) * (durations[j] * cycle_dist);

        var fly_steps = durations[j];
        for (0..(time_steps % cycle_time)) |_| {
            if (fly_steps > 0) {
                dists[j] += speeds[j];
                fly_steps -= 1;
            } else {
                break;
            }
        }
    }

    var ans: u32 = std.math.minInt(u32);
    for (0..i) |j| {
        ans = @max(ans, dists[j]);
    }

    return ans;
}

test "a__core" {
    const input_text =
        \\Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
        \\Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
    ;
    const allocator = std.testing.allocator;
    try std.testing.expectEqual(1120, try a__core(input_text, 1000, allocator));
}

pub fn a(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const winning_dist = try a__core(input_text, 2503, allocator);

    if (!config.benchmark) {
        std.debug.print("day 14 a: {d}\n", .{winning_dist});
    }
}

fn b__core(input_text: []const u8, time_steps: u32, _: std.mem.Allocator) !u32 {
    const NUM_REINDEER = 9;
    const MOVING_STATE: u32 = 1;

    var speeds = [_]u32{0} ** NUM_REINDEER;
    var durations = [_]u32{0} ** NUM_REINDEER;
    var rests = [_]u32{0} ** NUM_REINDEER;
    var states = [_]u1{0} ** NUM_REINDEER;
    var timers = [_]u32{0} ** NUM_REINDEER;
    var i: usize = 0;

    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.split(u8, line, " ");
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        const speed_raw = toks.next().?;
        const speed = try std.fmt.parseInt(u8, speed_raw, 10);
        _ = toks.next().?;
        _ = toks.next().?;
        const duration_raw = toks.next().?;
        const duration = try std.fmt.parseInt(u8, duration_raw, 10);
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        const rest_raw = toks.next().?;
        const rest = try std.fmt.parseInt(u8, rest_raw, 10);

        speeds[i] = speed;
        durations[i] = duration;
        rests[i] = rest;
        states[i] = MOVING_STATE;
        timers[i] = duration;
        i += 1;
    }

    var dists = [_]u32{0} ** NUM_REINDEER;
    var points = [_]u32{0} ** NUM_REINDEER;

    for (0..time_steps) |_| {
        var max_dist: u32 = std.math.minInt(u32);

        // Step all reindeer
        for (0..i) |j| {
            timers[j] -= 1;

            if (states[j] == MOVING_STATE) {
                dists[j] += speeds[j];
            }

            max_dist = @max(max_dist, dists[j]);

            if (timers[j] == 0) {
                states[j] = 1 - states[j];
                timers[j] = if (states[j] == 1) durations[j] else rests[j];
            }
        }

        // All players at the max distance get a point
        for (0..i) |j| {
            if (dists[j] == max_dist) {
                points[j] += 1;
            }
        }
    }

    var ans: u32 = std.math.minInt(u32);
    for (0..i) |j| {
        ans = @max(ans, points[j]);
    }

    return ans;
}

test "b__core" {
    const input_text =
        \\Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
        \\Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
    ;
    const allocator = std.testing.allocator;
    try std.testing.expectEqual(689, try b__core(input_text, 1000, allocator));
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const winning_dist = try b__core(input_text, 2503, allocator);

    if (!config.benchmark) {
        std.debug.print("day 14 b: {d}\n", .{winning_dist});
    }
}
