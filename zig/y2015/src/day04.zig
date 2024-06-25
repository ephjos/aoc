const std = @import("std");
const config = @import("config");

pub fn a(input_text: []const u8, _: std.mem.Allocator) !void {
    var lines = std.mem.splitAny(u8, input_text, "\n");
    const key = lines.first();

    var buf = [_]u8{0} ** std.crypto.hash.Md5.block_length;
    var digest_buf = [_]u8{0} ** std.crypto.hash.Md5.digest_length;

    var first_suffix_i: usize = 0;
    for (1..(1 << 32)) |i| {
        const coin = try std.fmt.bufPrint(&buf, "{s}{d}", .{ key, i });

        std.crypto.hash.Md5.hash(coin, &digest_buf, .{});

        if ((digest_buf[0] == 0x00) and (digest_buf[1] == 0x00) and (digest_buf[2] <= 0x0F)) {
            first_suffix_i = i;
            break;
        }
    }

    if (!config.benchmark) {
        std.debug.print("day  4 a: {d}\n", .{first_suffix_i});
    }
}

const ORDER = std.builtin.AtomicOrder.acquire;

fn checkMd5Interval(out: *std.atomic.Value(usize), key: []const u8, start: usize, end: usize) void {
    var buf = [_]u8{0} ** std.crypto.hash.Md5.block_length;
    var digest_buf = [_]u8{0} ** std.crypto.hash.Md5.digest_length;

    for (start..end) |i| {
        const coin = std.fmt.bufPrint(&buf, "{s}{d}", .{ key, i }) catch continue;

        std.crypto.hash.Md5.hash(coin, &digest_buf, .{});

        if ((digest_buf[0] == 0x00) and (digest_buf[1] == 0x00) and (digest_buf[2] == 0x00)) {
            _ = out.fetchMin(i, ORDER);
            return;
        }
    }
}

pub fn b(input_text: []const u8, allocator: std.mem.Allocator) !void {
    const INITIAL_SUFFIX = std.math.maxInt(usize);

    var thread_pool: std.Thread.Pool = undefined;
    try thread_pool.init(std.Thread.Pool.Options{
        .allocator = allocator,
    });
    defer thread_pool.deinit();

    var wait_group: std.Thread.WaitGroup = undefined;

    var first_suffix = std.atomic.Value(usize).init(INITIAL_SUFFIX);

    var lines = std.mem.splitAny(u8, input_text, "\n");
    const key = lines.first();

    // Sliding window that is broken up into blocks per thread
    const day04_window_size: usize = config.day04_window_size;
    var window_end: usize = day04_window_size;

    var block_start: usize = 0;
    const day04_block_size: usize = config.day04_block_size;

    while (first_suffix.load(ORDER) == INITIAL_SUFFIX) {
        wait_group.reset();

        while (block_start < window_end) : (block_start += day04_block_size) {
            thread_pool.spawnWg(&wait_group, checkMd5Interval, .{ &first_suffix, key, block_start, block_start + day04_block_size });
        }

        thread_pool.waitAndWork(&wait_group);
        window_end += day04_window_size;
    }

    if (!config.benchmark) {
        std.debug.print("day  4 b: {d}\n", .{first_suffix.load(ORDER)});
    }
}
