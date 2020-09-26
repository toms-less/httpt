load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "s2",
    srcs = [
        "cpuid_amd64.go",
        "cpuid_amd64.s",
        "decode.go",
        "decode_amd64.go",
        "decode_amd64.s",
        "decode_other.go",
        "encode.go",
        "encode_all.go",
        "encode_amd64.go",
        "encode_better.go",
        "encode_go.go",
        "encodeblock_amd64.go",
        "encodeblock_amd64.s",
        "s2.go",
    ],
    importpath = "github.com/klauspost/compress/s2",
    visibility = ["//visibility:public"],
)

go_test(
    name = "s2_test",
    srcs = [
        "decode_test.go",
        "encode_test.go",
        "s2_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":s2"],
    deps = [
        "//deps/github.com/klauspost/compress/snappy",
        "//deps/github.com/klauspost/compress/zip",
        "//deps/github.com/klauspost/compress/zstd",
    ],
)
