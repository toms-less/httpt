load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "file",
    srcs = [
        "downloader.go",
        "file.go",
        "reader.go",
        "writer.go",
    ],
    importpath = "github.com/toolkits/file",
    visibility = ["//visibility:public"],
)

go_test(
    name = "file_test",
    srcs = ["file_test.go"],
    embed = [":file"],
)
