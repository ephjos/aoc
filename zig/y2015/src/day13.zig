const std = @import("std");
const config = @import("config");

// Modified Held-Karp to find max:
// https://en.wikipedia.org/wiki/Held%E2%80%93Karp_algorithm
fn hk(n: usize, d: []const i32, allocator: std.mem.Allocator) !i32 {
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

pub fn a__core(input_text: []const u8, allocator: std.mem.Allocator) !i32 {
    var index: usize = 0;
    var n2i = std.StringHashMap(usize).init(allocator);
    defer n2i.deinit();

    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.splitAny(u8, line, " ");
        const n0 = toks.first();
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        const n1 = std.mem.trim(u8, toks.next().?, ". \n");

        if (!n2i.contains(n0)) {
            try n2i.put(n0, index);
            index += 1;
        }

        if (!n2i.contains(n1)) {
            try n2i.put(n1, index);
            index += 1;
        }
    }

    const V = index;
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
        const n0 = toks.first();
        _ = toks.next().?;
        const sign_r = toks.next().?;
        const delta_r = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        const n1 = std.mem.trim(u8, toks.next().?, ". \n");

        const sign: i32 = if (std.mem.eql(u8, "lose", sign_r)) -1 else 1;
        const delta = try std.fmt.parseInt(u8, delta_r, 10);

        const n0i = n2i.get(n0).?;
        const n1i = n2i.get(n1).?;

        dist[n0i * V + n1i] = sign * delta;

        dist[n0i * V + n0i] = 0;
        dist[n1i * V + n1i] = 0;
    }

    // Make dist symmetrical:
    // dist[i][j] = dist[j][i] = total happiness placing i next to j
    for (0..V) |i| {
        for (i..V) |j| {
            const v = dist[i * V + j] + dist[j * V + i];
            dist[i * V + j] = v;
            dist[j * V + i] = v;
        }
    }

    //    for (0..V) |i| {
    //        for (0..V) |j| {
    //            const x = dist[i * V + j];
    //            if (x == std.math.minInt(i32)) {
    //                std.debug.print("  -inf ", .{});
    //            } else {
    //                std.debug.print("{d:6} ", .{dist[i * V + j]});
    //            }
    //        }
    //        std.debug.print("\n", .{});
    //    }

    return try hk(V, dist, allocator);
}

test "a__core" {
    const allocator = std.testing.allocator;
    const input_text =
        \\Alice would gain 54 happiness units by sitting next to Bob.
        \\Alice would lose 79 happiness units by sitting next to Carol.
        \\Alice would lose 2 happiness units by sitting next to David.
        \\Bob would gain 83 happiness units by sitting next to Alice.
        \\Bob would lose 7 happiness units by sitting next to Carol.
        \\Bob would lose 63 happiness units by sitting next to David.
        \\Carol would lose 62 happiness units by sitting next to Alice.
        \\Carol would gain 60 happiness units by sitting next to Bob.
        \\Carol would gain 55 happiness units by sitting next to David.
        \\David would gain 46 happiness units by sitting next to Alice.
        \\David would lose 7 happiness units by sitting next to Bob.
        \\David would gain 41 happiness units by sitting next to Carol.
    ;

    try std.testing.expectEqual(330, try a__core(input_text, allocator));
}

pub fn a(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const answer = try a__core(input_text, allocator);

    if (!config.benchmark) {
        std.debug.print("day 13 a: {d}\n", .{answer});
    }
}

pub fn b__core(input_text: []const u8, allocator: std.mem.Allocator) !i32 {
    var index: usize = 1;
    var n2i = std.StringHashMap(usize).init(allocator);
    defer n2i.deinit();

    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.splitAny(u8, line, " ");
        const n0 = toks.first();
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        const n1 = std.mem.trim(u8, toks.next().?, ". \n");

        if (!n2i.contains(n0)) {
            try n2i.put(n0, index);
            index += 1;
        }

        if (!n2i.contains(n1)) {
            try n2i.put(n1, index);
            index += 1;
        }
    }

    const V = index;
    const dist = try allocator.alloc(i32, V * V);
    defer allocator.free(dist);
    for (0..V * V) |i| {
        dist[i] = 0;
    }

    for (0..V) |i| {
        dist[i * V + (V - 1)] = 0;
        dist[(V - 1) * V + i] = 0;
    }

    lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.splitAny(u8, line, " ");
        const n0 = toks.first();
        _ = toks.next().?;
        const sign_r = toks.next().?;
        const delta_r = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        _ = toks.next().?;
        const n1 = std.mem.trim(u8, toks.next().?, ". \n");

        const sign: i32 = if (std.mem.eql(u8, "lose", sign_r)) -1 else 1;
        const delta = try std.fmt.parseInt(u8, delta_r, 10);

        const n0i = n2i.get(n0).?;
        const n1i = n2i.get(n1).?;

        dist[n0i * V + n1i] = sign * delta;

        dist[n0i * V + n0i] = 0;
        dist[n1i * V + n1i] = 0;
    }

    // Make dist symmetrical:
    // dist[i][j] = dist[j][i] = total happiness placing i next to j
    for (0..V) |i| {
        for (i..V) |j| {
            const v = dist[i * V + j] + dist[j * V + i];
            dist[i * V + j] = v;
            dist[j * V + i] = v;
        }
    }

    //    for (0..V) |i| {
    //        for (0..V) |j| {
    //            const x = dist[i * V + j];
    //            if (x == std.math.minInt(i32)) {
    //                std.debug.print("  -inf ", .{});
    //            } else {
    //                std.debug.print("{d:6} ", .{dist[i * V + j]});
    //            }
    //        }
    //        std.debug.print("\n", .{});
    //    }
    return try hk(V, dist, allocator);
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const answer = try b__core(input_text, allocator);

    if (!config.benchmark) {
        std.debug.print("day 13 b: {d}\n", .{answer});
    }
}
