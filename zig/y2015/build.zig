const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    // Define options
    const benchmark = b.option(bool, "benchmark", "Build in benchmark mode") orelse false;
    const samples = b.option(u32, "samples", "Number of iterations to time in benchmark mode") orelse 8;

    const options = b.addOptions();
    options.addOption(bool, "benchmark", benchmark);
    options.addOption(u32, "samples", samples);

    // Setup `zig build test` mem check
    const test_step = b.step("test", "Run unit tests");
    const unit_tests = b.addTest(.{
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });
    unit_tests.root_module.addOptions("config", options);
    const run_unit_tests = b.addRunArtifact(unit_tests);
    test_step.dependOn(&run_unit_tests.step);

    // Define main executable
    const exe = b.addExecutable(.{
        .name = "y2015",
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });
    exe.root_module.addOptions("config", options);

    b.installArtifact(exe);

    // Define `zig build run` step
    const run_exe = b.addRunArtifact(exe);

    const run_step = b.step("run", "Run the application");
    run_step.dependOn(&run_exe.step);
}
