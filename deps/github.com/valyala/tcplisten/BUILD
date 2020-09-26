load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "tcplisten",
    srcs = [
        "socket.go",
        "socket_darwin.go",
        "socket_other.go",
        "tcplisten.go",
        "tcplisten_bsd.go",
        "tcplisten_linux.go",
    ],
    importpath = "github.com/valyala/tcplisten",
    visibility = ["//visibility:public"],
)

go_test(
    name = "tcplisten_test",
    srcs = ["tcplisten_test.go"],
    embed = [":tcplisten"],
)
