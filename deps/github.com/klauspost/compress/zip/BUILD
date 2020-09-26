load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "zip",
    srcs = [
        "reader.go",
        "register.go",
        "struct.go",
        "writer.go",
    ],
    importpath = "github.com/klauspost/compress/zip",
    visibility = ["//visibility:public"],
    deps = ["//deps/github.com/klauspost/compress/flate"],
)

go_test(
    name = "zip_test",
    srcs = [
        "example_test.go",
        "reader_test.go",
        "writer_test.go",
        "zip_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":zip"],
    deps = ["//deps/github.com/klauspost/compress/flate"],
)
