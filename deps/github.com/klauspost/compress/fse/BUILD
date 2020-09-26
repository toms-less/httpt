load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "fse",
    srcs = [
        "bitreader.go",
        "bitwriter.go",
        "bytereader.go",
        "compress.go",
        "decompress.go",
        "fse.go",
    ],
    importpath = "github.com/klauspost/compress/fse",
    visibility = ["//visibility:public"],
)

go_test(
    name = "fse_test",
    srcs = ["fse_test.go"],
    embed = [":fse"],
)
