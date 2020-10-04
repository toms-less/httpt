load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "bytebufferpool",
    srcs = [
        "bytebuffer.go",
        "doc.go",
        "pool.go",
    ],
    importpath = "github.com/valyala/bytebufferpool",
    visibility = ["//visibility:public"],
)

go_test(
    name = "bytebufferpool_test",
    srcs = [
        "bytebuffer_example_test.go",
        "bytebuffer_test.go",
        "bytebuffer_timing_test.go",
        "pool_test.go",
    ],
    embed = [":bytebufferpool"],
)
