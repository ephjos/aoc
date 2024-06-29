const std = @import("std");
const config = @import("config");

// Held-Karp:
// https://en.wikipedia.org/wiki/Held%E2%80%93Karp_algorithm
fn a__get_shortest(n: usize, d: []const i32, allocator: std.mem.Allocator) !i32 {
    const bits_space = @as(u32, 1) << @intCast(n);

    // Init G
    var G = try allocator.alloc(i32, n * bits_space);
    defer allocator.free(G);
    for (0..G.len) |i| {
        G[i] = std.math.maxInt(i32);
    }
    for (1..n) |k| {
        G[(k * bits_space) + (@as(u32, 1) << @intCast(k))] = d[(0 * n) + k];
    }

    var sets = std.ArrayList(u32).init(allocator);
    defer sets.deinit();
    for (2..n) |s| {
        // Build sets of size s
        sets.clearAndFree();
        for (1..bits_space) |j| {
            if (@popCount(j) == s) {
                try sets.append(@intCast(j));
            }
        }

        // Core loop
        for (sets.items) |S| {
            for (1..n) |k| {
                const k_bit = @as(u32, 1) << @intCast(k);
                if (S & k_bit == 0) {
                    continue;
                }

                var opt: i32 = std.math.maxInt(i32);
                const rest = S & ~k_bit;
                for (1..n) |m| {
                    const m_bit = @as(u32, 1) << @intCast(m);
                    if (S & m_bit == 0 or m == k or m == 0) {
                        continue;
                    }
                    const x = G[(m * bits_space) + rest];
                    const y = d[(m * n) + k];
                    if (x == std.math.maxInt(i32) or y == std.math.maxInt(i32)) {
                        continue;
                    }
                    opt = @min(opt, x + y);
                }
                G[(k * bits_space) + S] = opt;
            }
        }
    }

    var opt: i32 = std.math.maxInt(i32);
    const rest = bits_space - 2;
    for (1..n) |k| {
        const x = G[(k * bits_space) + rest];
        const y = d[(k * n) + 0];
        opt = @min(opt, x + y);
    }

    return opt;
}

fn find_shortest(input_text: []const u8, allocator: std.mem.Allocator) !i32 {
    var index: usize = 1;
    var city_to_index = std.hash_map.StringHashMap(usize).init(allocator);
    defer city_to_index.deinit();

    // Find all city_to_index
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.splitAny(u8, line, " ");
        const c1 = toks.first();
        _ = toks.next().?; // to
        const c2 = toks.next().?;

        if (!city_to_index.contains(c1)) {
            try city_to_index.put(c1, index);
            index += 1;
        }

        if (!city_to_index.contains(c2)) {
            try city_to_index.put(c2, index);
            index += 1;
        }
    }

    // Init dist
    const V = city_to_index.count() + 1;
    const dist = try allocator.alloc(i32, V * V);
    defer allocator.free(dist);
    for (0..V * V) |i| {
        dist[i] = 0;
    }

    lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.splitAny(u8, line, " ");
        const c1 = toks.first();
        _ = toks.next().?; // to
        const c2 = toks.next().?;
        _ = toks.next().?; // =
        const weight = toks.next().?;

        const c1_i = city_to_index.get(c1).?;
        const c2_i = city_to_index.get(c2).?;
        const weight_int = try std.fmt.parseInt(i32, weight, 10);

        dist[c1_i * V + c2_i] = weight_int;
        dist[c2_i * V + c1_i] = weight_int;

        dist[c1_i * V + c1_i] = 0;
        dist[c2_i * V + c2_i] = 0;
    }

    return try a__get_shortest(V, dist, allocator);
}

test "find_shortest" {
    const input_text =
        \\London to Dublin = 464
        \\London to Belfast = 518
        \\Dublin to Belfast = 141
    ;
    const allocator = std.testing.allocator;

    try std.testing.expectEqual(605, try find_shortest(input_text, allocator));
}

pub fn a(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const length = try find_shortest(input_text, allocator);

    if (!config.benchmark) {
        std.debug.print("day  9 a: {d}\n", .{length});
    }
}

// Modified Held-Karp to find max:
// https://en.wikipedia.org/wiki/Held%E2%80%93Karp_algorithm
fn b__get_longest(n: usize, d: []const i32, allocator: std.mem.Allocator) !i32 {
    const bits_space = @as(u32, 1) << @intCast(n);

    // Init G
    var G = try allocator.alloc(i32, n * bits_space);
    defer allocator.free(G);
    for (0..G.len) |i| {
        G[i] = std.math.minInt(i32);
    }
    for (1..n) |k| {
        G[(k * bits_space) + (@as(u32, 1) << @intCast(k))] = d[(0 * n) + k];
    }

    var sets = std.ArrayList(u32).init(allocator);
    defer sets.deinit();
    for (2..n) |s| {
        // Build sets of size s
        sets.clearAndFree();
        for (1..bits_space) |j| {
            if (@popCount(j) == s) {
                try sets.append(@intCast(j));
            }
        }

        // Core loop
        for (sets.items) |S| {
            for (1..n) |k| {
                const k_bit = @as(u32, 1) << @intCast(k);
                if (S & k_bit == 0) {
                    continue;
                }

                var opt: i32 = std.math.minInt(i32);
                const rest = S & ~k_bit;
                for (1..n) |m| {
                    const m_bit = @as(u32, 1) << @intCast(m);
                    if (S & m_bit == 0 or m == k or m == 0) {
                        continue;
                    }
                    const x = G[(m * bits_space) + rest];
                    const y = d[(m * n) + k];
                    if (x == std.math.minInt(i32) or y == std.math.minInt(i32)) {
                        continue;
                    }
                    opt = @max(opt, x + y);
                }
                G[(k * bits_space) + S] = opt;
            }
        }
    }

    var opt: i32 = std.math.minInt(i32);
    const rest = bits_space - 2;
    for (1..n) |k| {
        const x = G[(k * bits_space) + rest];
        const y = d[(k * n) + 0];
        opt = @max(opt, x + y);
    }

    return opt;
}
fn find_longest(input_text: []const u8, allocator: std.mem.Allocator) !i32 {
    var index: usize = 1;
    var city_to_index = std.hash_map.StringHashMap(usize).init(allocator);
    defer city_to_index.deinit();

    // Find all city_to_index
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.splitAny(u8, line, " ");
        const c1 = toks.first();
        _ = toks.next().?; // to
        const c2 = toks.next().?;

        if (!city_to_index.contains(c1)) {
            try city_to_index.put(c1, index);
            index += 1;
        }

        if (!city_to_index.contains(c2)) {
            try city_to_index.put(c2, index);
            index += 1;
        }
    }

    // Init dist
    const V = city_to_index.count() + 1;
    const dist = try allocator.alloc(i32, V * V);
    defer allocator.free(dist);
    for (0..V * V) |i| {
        dist[i] = 0;
    }

    lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.splitAny(u8, line, " ");
        const c1 = toks.first();
        _ = toks.next().?; // to
        const c2 = toks.next().?;
        _ = toks.next().?; // =
        const weight = toks.next().?;

        const c1_i = city_to_index.get(c1).?;
        const c2_i = city_to_index.get(c2).?;
        const weight_int = try std.fmt.parseInt(i32, weight, 10);

        dist[c1_i * V + c2_i] = weight_int;
        dist[c2_i * V + c1_i] = weight_int;

        dist[c1_i * V + c1_i] = 0;
        dist[c2_i * V + c2_i] = 0;
    }

    return try b__get_longest(V, dist, allocator);
}

test "find_longest" {
    const input_text =
        \\London to Dublin = 464
        \\London to Belfast = 518
        \\Dublin to Belfast = 141
    ;
    const allocator = std.testing.allocator;

    try std.testing.expectEqual(982, try find_longest(input_text, allocator));
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const length = try find_longest(input_text, allocator);

    if (!config.benchmark) {
        std.debug.print("day  9 b: {d}\n", .{length});
    }
}
