load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "log4go",
    srcs = [
        "category.go",
        "filelog.go",
        "jsonconfig.go",
        "log4go.go",
        "pattlog.go",
        "socklog.go",
        "termlog.go",
        "util.go",
        "wrapper.go",
        "xmlconfig.go",
    ],
    importpath = "github.com/jeanphorn/log4go",
    visibility = ["//visibility:public"],
    deps = ["//deps/github.com/toolkits/file"],
)
