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

fn find_shortest(input_text: []const u8, allocator: std.mem.Allocator) !u32 {
    var index: usize = 0;
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
    const V = city_to_index.count();
    const dist = try allocator.alloc(u32, V * V);
    defer allocator.free(dist);
    for (0..V * V) |i| {
        dist[i] = std.math.maxInt(u32);
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
        const weight_int = try std.fmt.parseInt(u32, weight, 10);

        dist[c1_i * V + c2_i] = weight_int;
        dist[c2_i * V + c1_i] = weight_int;

        dist[c1_i * V + c1_i] = 0;
        dist[c2_i * V + c2_i] = 0;
    }

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

    var answer: u32 = std.math.maxInt(u32);
    for (permutations.items) |p| {
        var cost: u32 = 0;
        var previous = p[0];
        for (p[1..]) |node| {
            cost += dist[previous * V + node];
            previous = node;
        }
        answer = @min(answer, cost);
    }

    return answer;
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

fn find_longest(input_text: []const u8, allocator: std.mem.Allocator) !u32 {
    var index: usize = 0;
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
    const V = city_to_index.count();
    const dist = try allocator.alloc(u32, V * V);
    defer allocator.free(dist);
    for (0..V * V) |i| {
        dist[i] = std.math.maxInt(u32);
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
        const weight_int = try std.fmt.parseInt(u32, weight, 10);

        dist[c1_i * V + c2_i] = weight_int;
        dist[c2_i * V + c1_i] = weight_int;

        dist[c1_i * V + c1_i] = 0;
        dist[c2_i * V + c2_i] = 0;
    }

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

    var answer: u32 = 0;
    for (permutations.items) |p| {
        var cost: u32 = 0;
        var previous = p[0];
        for (p[1..]) |node| {
            cost += dist[previous * V + node];
            previous = node;
        }
        answer = @max(answer, cost);
    }

    return answer;
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
