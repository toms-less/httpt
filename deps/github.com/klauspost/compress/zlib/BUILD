load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "zlib",
    srcs = [
        "reader.go",
        "writer.go",
    ],
    importpath = "github.com/klauspost/compress/zlib",
    visibility = ["//visibility:public"],
    deps = ["//deps/github.com/klauspost/compress/flate"],
)

go_test(
    name = "zlib_test",
    srcs = [
        "example_test.go",
        "reader_test.go",
        "writer_test.go",
    ],
    embed = [":zlib"],
)
