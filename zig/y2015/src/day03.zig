const std = @import("std");
const config = @import("config");
const input_text = @embedFile("./inputs/d03");

const Point = @Vector(2, i32);

const UP = Point{ 0, 1 };
const DOWN = Point{ 0, -1 };
const LEFT = Point{ -1, 0 };
const RIGHT = Point{ 1, 0 };

pub fn a(allocator: std.mem.Allocator) anyerror!void {
    var houses = std.hash_map.AutoHashMap(Point, void).init(allocator);
    defer houses.deinit();

    var curr = Point{ 0, 0 };
    try houses.put(curr, {});

    for (input_text) |c| {
        switch (c) {
            '^' => curr += UP,
            'v' => curr += DOWN,
            '<' => curr += LEFT,
            '>' => curr += RIGHT,
            else => {},
        }

        try houses.put(curr, {});
    }

    if (!config.benchmark) {
        std.debug.print("day  3 a: {d}\n", .{houses.count()});
    }
}

pub fn b(allocator: std.mem.Allocator) anyerror!void {
    var houses = std.hash_map.AutoHashMap(Point, void).init(allocator);
    defer houses.deinit();

    var santa = Point{ 0, 0 };
    var robo = Point{ 0, 0 };
    try houses.put(santa, {});
    try houses.put(robo, {});

    for (input_text, 0..) |c, i| {
        if (i % 2 == 0) {
            switch (c) {
                '^' => santa += UP,
                'v' => santa += DOWN,
                '<' => santa += LEFT,
                '>' => santa += RIGHT,
                else => {},
            }
            try houses.put(santa, {});
        } else {
            switch (c) {
                '^' => robo += UP,
                'v' => robo += DOWN,
                '<' => robo += LEFT,
                '>' => robo += RIGHT,
                else => {},
            }
            try houses.put(robo, {});
        }
    }

    if (!config.benchmark) {
        std.debug.print("day  3 b: {d}\n", .{houses.count()});
    }
}
