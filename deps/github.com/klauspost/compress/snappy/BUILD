load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "snappy",
    srcs = [
        "decode.go",
        "decode_amd64.go",
        "decode_amd64.s",
        "decode_other.go",
        "encode.go",
        "encode_amd64.go",
        "encode_amd64.s",
        "encode_other.go",
        "snappy.go",
    ],
    importpath = "github.com/klauspost/compress/snappy",
    visibility = ["//visibility:public"],
)

go_test(
    name = "snappy_test",
    srcs = [
        "golden_test.go",
        "snappy_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":snappy"],
)
