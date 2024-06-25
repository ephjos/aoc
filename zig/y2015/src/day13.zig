const std = @import("std");
const config = @import("config");

fn generate(k: u32, A: []usize, output: *std.ArrayList([]usize), allocator: std.mem.Allocator) !void {
    if (k == 1) {
        try output.append(try allocator.dupe(usize, A));
        return;
    }

    for (0..k) |i| {
        try generate(k - 1, A, output, allocator);
        if (k % 2 == 0) {
            const temp = A[i];
            A[i] = A[k - 1];
            A[k - 1] = temp;
        } else {
            const temp = A[0];
            A[0] = A[k - 1];
            A[k - 1] = temp;
        }
    }
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

    const V = n2i.count();
    const dist = try allocator.alloc(i32, V * V);
    defer allocator.free(dist);
    for (0..V * V) |i| {
        dist[i] = std.math.minInt(i32);
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

    // Generate and check every permutation of possible paths
    const A = try allocator.alloc(usize, V);
    defer allocator.free(A);
    for (0..V) |i| {
        A[i] = i;
    }
    var permutations = std.ArrayList([]usize).init(allocator);
    defer {
        for (permutations.items) |p| {
            allocator.free(p);
        }
        permutations.deinit();
    }

    try generate(V, A, &permutations, allocator);

    var answer: i32 = std.math.minInt(i32);
    for (permutations.items) |p| {
        var happiness: i32 = 0;
        var previous = p[0];
        for (p[1..]) |node| {
            happiness += dist[previous * V + node];
            happiness += dist[node * V + previous];
            previous = node;
        }
        happiness += dist[previous * V + p[0]];
        happiness += dist[p[0] * V + previous];

        answer = @max(answer, happiness);
    }

    return answer;
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

    const V = n2i.count() + 1;
    const dist = try allocator.alloc(i32, V * V);
    defer allocator.free(dist);
    for (0..V * V) |i| {
        dist[i] = std.math.minInt(i32);
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

    // Generate and check every permutation of possible paths
    const A = try allocator.alloc(usize, V);
    defer allocator.free(A);
    for (0..V) |i| {
        A[i] = i;
    }
    var permutations = std.ArrayList([]usize).init(allocator);
    defer {
        for (permutations.items) |p| {
            allocator.free(p);
        }
        permutations.deinit();
    }

    try generate(V, A, &permutations, allocator);

    var answer: i32 = std.math.minInt(i32);
    for (permutations.items) |p| {
        var happiness: i32 = 0;
        var previous = p[0];
        for (p[1..]) |node| {
            happiness += dist[previous * V + node];
            happiness += dist[node * V + previous];
            previous = node;
        }
        happiness += dist[previous * V + p[0]];
        happiness += dist[p[0] * V + previous];

        answer = @max(answer, happiness);
    }

    return answer;
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const answer = try b__core(input_text, allocator);

    if (!config.benchmark) {
        std.debug.print("day 13 b: {d}\n", .{answer});
    }
}
