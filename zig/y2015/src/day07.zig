const std = @import("std");
const config = @import("config");

const InstructionType = enum(u16) { SET, AND, OR, NOT, LSHIFT, RSHIFT };

const Instruction = union(InstructionType) {
    SET: struct {
        imm: []const u8,
    },
    AND: struct {
        lhs: []const u8,
        rhs: []const u8,
    },
    OR: struct {
        lhs: []const u8,
        rhs: []const u8,
    },
    NOT: struct {
        src: []const u8,
    },
    LSHIFT: struct {
        lhs: []const u8,
        rhs: []const u8,
    },
    RSHIFT: struct {
        lhs: []const u8,
        rhs: []const u8,
    },
};

const Memory = std.hash_map.StringHashMap(Instruction);
const EvalMemory = std.hash_map.StringHashMap(u16);

fn load_mem(mem: *Memory, input_text: []const u8) !void {
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var toks = std.mem.split(u8, line, " -> ");
        const lhs = toks.first();
        const rhs = toks.next().?;

        var lhs_toks = std.mem.split(u8, lhs, " ");
        const t0 = lhs_toks.first();
        const t1_r = lhs_toks.next();

        if (t1_r == null) {
            try mem.put(rhs, .{ .SET = .{ .imm = t0 } });
            continue;
        }

        const t1 = t1_r.?;

        if (std.mem.eql(u8, t0, "NOT") or std.mem.eql(u8, t1, "NOT")) {
            try mem.put(rhs, .{ .NOT = .{ .src = t1 } });
            continue;
        }

        const t2 = lhs_toks.next().?;

        if (std.mem.eql(u8, t1, "AND")) {
            try mem.put(rhs, .{ .AND = .{ .lhs = t0, .rhs = t2 } });
            continue;
        }

        if (std.mem.eql(u8, t1, "OR")) {
            try mem.put(rhs, .{ .OR = .{ .lhs = t0, .rhs = t2 } });
            continue;
        }

        if (std.mem.eql(u8, t1, "LSHIFT")) {
            try mem.put(rhs, .{ .LSHIFT = .{ .lhs = t0, .rhs = t2 } });
            continue;
        }

        if (std.mem.eql(u8, t1, "RSHIFT")) {
            try mem.put(rhs, .{ .RSHIFT = .{ .lhs = t0, .rhs = t2 } });
            continue;
        }

        std.debug.print("Should be unreachable, something went wrong with: {s}\n", .{line});
    }

    return;
}

fn eval(mem: *Memory, eval_mem: *EvalMemory, reg: []const u8) !u16 {
    if (std.ascii.isDigit(reg[0])) {
        return std.fmt.parseInt(u16, reg, 10);
    }

    const eval_mem_val = eval_mem.get(reg);
    if (eval_mem_val != null) {
        return eval_mem_val.?;
    }

    const instr = mem.get(reg).?;

    // TODO: set eval_mem
    const foo = switch (instr) {
        .SET => try eval(mem, eval_mem, instr.SET.imm),
        .AND => try eval(mem, eval_mem, instr.AND.lhs) & try eval(mem, eval_mem, instr.AND.rhs),
        .OR => try eval(mem, eval_mem, instr.OR.lhs) | try eval(mem, eval_mem, instr.OR.rhs),
        .NOT => ~try eval(mem, eval_mem, instr.NOT.src),
        .LSHIFT => try eval(mem, eval_mem, instr.LSHIFT.lhs) << @intCast(try eval(mem, eval_mem, instr.LSHIFT.rhs)),
        .RSHIFT => (try eval(mem, eval_mem, instr.RSHIFT.lhs) >> @intCast(try eval(mem, eval_mem, instr.RSHIFT.rhs))),
    };
    try eval_mem.put(reg, foo);
    return foo;
}

test "load_mem" {
    var mem = Memory.init(std.testing.allocator);
    defer mem.deinit();
    var eval_mem = EvalMemory.init(std.testing.allocator);
    defer eval_mem.deinit();

    const input_text =
        \\123 -> x
        \\456 -> y
        \\x AND y -> d
        \\x OR y -> e
        \\x LSHIFT 2 -> f
        \\y RSHIFT 2 -> g
        \\NOT x -> h
        \\NOT y -> i
    ;

    try load_mem(&mem, input_text);

    try std.testing.expectEqual(72, try eval(&mem, &eval_mem, "d"));
    try std.testing.expectEqual(507, try eval(&mem, &eval_mem, "e"));
    try std.testing.expectEqual(492, try eval(&mem, &eval_mem, "f"));
    try std.testing.expectEqual(114, try eval(&mem, &eval_mem, "g"));
    try std.testing.expectEqual(65412, try eval(&mem, &eval_mem, "h"));
    try std.testing.expectEqual(65079, try eval(&mem, &eval_mem, "i"));
    try std.testing.expectEqual(123, try eval(&mem, &eval_mem, "x"));
    try std.testing.expectEqual(456, try eval(&mem, &eval_mem, "y"));
}

pub fn a(input_text: []const u8, allocator: std.mem.Allocator) !void {
    var mem = Memory.init(allocator);
    defer mem.deinit();
    var eval_mem = EvalMemory.init(allocator);
    defer eval_mem.deinit();

    try load_mem(&mem, input_text);

    const a_val = try eval(&mem, &eval_mem, "a");

    if (!config.benchmark) {
        std.debug.print("day  7 a: {d}\n", .{a_val});
    }
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const a_val = blk: {
        var mem = Memory.init(allocator);
        defer mem.deinit();
        var eval_mem = EvalMemory.init(allocator);
        defer eval_mem.deinit();

        try load_mem(&mem, input_text);
        break :blk try eval(&mem, &eval_mem, "a");
    };

    const b_val = blk: {
        var mem = Memory.init(allocator);
        defer mem.deinit();
        var eval_mem = EvalMemory.init(allocator);
        defer eval_mem.deinit();

        try load_mem(&mem, input_text);
        try eval_mem.put("b", a_val);
        break :blk try eval(&mem, &eval_mem, "a");
    };

    if (!config.benchmark) {
        std.debug.print("day  7 b: {d}\n", .{b_val});
    }
}
