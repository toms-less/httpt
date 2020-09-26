load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "compress",
    srcs = ["compressible.go"],
    importpath = "github.com/klauspost/compress",
    visibility = ["//visibility:public"],
)

go_test(
    name = "compress_test",
    srcs = ["compressible_test.go"],
    data = glob(["testdata/**"]),
    embed = [":compress"],
    deps = [
        "//deps/github.com/klauspost/compress/flate",
        "//deps/github.com/klauspost/compress/gzip",
    ],
)
