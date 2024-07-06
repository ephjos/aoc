const std = @import("std");
const config = @import("config");

pub fn a(input_text: []const u8, _: std.mem.Allocator) !void {
    var lines = std.mem.splitAny(u8, input_text, "\n");
    const key = lines.first();

    var input = [_]u8{0} ** std.crypto.hash.Md5.block_length;
    var digest_buf = [_]u8{0} ** std.crypto.hash.Md5.digest_length;

    const first_input = try std.fmt.bufPrint(&input, "{s}{d}", .{ key, 0 });
    var input_len = first_input.len;

    var first_suffix_i: usize = 0;
    for (1..(1 << 32)) |i| {
        if (input[input_len - 1] == '9') {
            var j = input_len - 1;
            while (input[j] == '9') {
                input[j] = '0';
                j -= 1;
            }

            if (j < key.len) {
                input_len += 1;
                input[input_len - 1] = '0';
                input[j + 1] = '1';
            } else {
                input[j] += 1;
            }
        } else {
            input[input_len - 1] += 1;
        }

        std.crypto.hash.Md5.hash(input[0..input_len], &digest_buf, .{});

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
    var input = [_]u8{0} ** std.crypto.hash.Md5.block_length;
    var digest_buf = [_]u8{0} ** std.crypto.hash.Md5.digest_length;

    const first_input = std.fmt.bufPrint(&input, "{s}{d}", .{ key, start - 1 }) catch unreachable;
    var input_len = first_input.len;

    for (start..end) |i| {
        if (input[input_len - 1] == '9') {
            var j = input_len - 1;
            while (input[j] == '9') {
                input[j] = '0';
                j -= 1;
            }

            if (j < key.len) {
                input_len += 1;
                input[input_len - 1] = '0';
                input[j + 1] = '1';
            } else {
                input[j] += 1;
            }
        } else {
            input[input_len - 1] += 1;
        }

        std.crypto.hash.Md5.hash(input[0..input_len], &digest_buf, .{});

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

    const num_threads = try std.Thread.getCpuCount();

    var block_start: usize = 1;
    const day04_block_size: usize = config.day04_block_size;

    // Sliding window that is broken up into blocks per thread
    const day04_window_size: usize = day04_block_size * num_threads;
    var window_end: usize = day04_window_size;

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
