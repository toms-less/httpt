load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "gzip",
    srcs = [
        "gunzip.go",
        "gzip.go",
    ],
    importpath = "github.com/klauspost/compress/gzip",
    visibility = ["//visibility:public"],
    deps = ["//deps/github.com/klauspost/compress/flate"],
)

go_test(
    name = "gzip_test",
    srcs = [
        "example_test.go",
        "gunzip_test.go",
        "gzip_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":gzip"],
    deps = ["//deps/github.com/klauspost/compress/flate"],
)
