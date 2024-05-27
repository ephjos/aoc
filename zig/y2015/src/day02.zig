const std = @import("std");
const config = @import("config");
const input_text = @embedFile("./inputs/d02");

pub fn a(_: std.mem.Allocator) anyerror!void {
    var total_paper_sqft: u32 = 0;
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var dims = std.mem.splitAny(u8, line, "x");
        const l = try std.fmt.parseInt(u32, dims.first(), 10);
        const w = try std.fmt.parseInt(u32, dims.next().?, 10);
        const h = try std.fmt.parseInt(u32, dims.next().?, 10);

        const lw = l * w;
        const wh = w * h;
        const lh = l * h;

        const slack = @min(lw, @min(wh, lh));

        total_paper_sqft += (2 * lw) + (2 * wh) + (2 * lh) + slack;
    }

    if (!config.benchmark) {
        std.debug.print("day  2 a: {d}\n", .{total_paper_sqft});
    }
}

pub fn b(_: std.mem.Allocator) anyerror!void {
    var total_ribbon_ft: u32 = 0;
    var lines = std.mem.splitAny(u8, input_text, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var dims = std.mem.splitAny(u8, line, "x");
        const l = try std.fmt.parseInt(u32, dims.first(), 10);
        const w = try std.fmt.parseInt(u32, dims.next().?, 10);
        const h = try std.fmt.parseInt(u32, dims.next().?, 10);

        const lw = l * w;
        const wh = w * h;
        const lh = l * h;
        const vol = l * w * h;

        const slack = @min(lw, @min(wh, lh));

        if (slack == lw) {
            total_ribbon_ft += (2 * l) + (2 * w) + vol;
        } else if (slack == wh) {
            total_ribbon_ft += (2 * w) + (2 * h) + vol;
        } else if (slack == lh) {
            total_ribbon_ft += (2 * l) + (2 * h) + vol;
        }
    }

    if (!config.benchmark) {
        std.debug.print("day  2 b: {d}\n", .{total_ribbon_ft});
    }
}
