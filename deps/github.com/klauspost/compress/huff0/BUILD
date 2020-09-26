load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "huff0",
    srcs = [
        "bitreader.go",
        "bitwriter.go",
        "bytereader.go",
        "compress.go",
        "decompress.go",
        "huff0.go",
    ],
    importpath = "github.com/klauspost/compress/huff0",
    visibility = ["//visibility:public"],
    deps = ["//deps/github.com/klauspost/compress/fse"],
)

go_test(
    name = "huff0_test",
    srcs = [
        "compress_test.go",
        "decompress_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":huff0"],
    deps = [
        "//deps/github.com/klauspost/compress/flate",
        "//deps/github.com/klauspost/compress/zip",
    ],
)
