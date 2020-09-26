package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")
load("//:build/workspace.bzl", "RELEASE_VERSION")

#build package of dashboard.
pkg_tar(
    name = "httpt-" + RELEASE_VERSION,
    srcs = [":httpt"],
    extension = "tar.gz",
    mode = "0755",
)

go_binary(
    name = "httpt",
    embed = [":main"],
)

go_library(
    name = "main",
    srcs = ["app.go"],
    importpath = "github.com/toms-less/httpt",
    deps = [
        ":pkg_config",
        ":pkg_logger",
        ":pkg_server",
        ":pkg_uri",
    ],
)

go_library(
    name = "pkg_caller",
    srcs = [
        "pkg/caller/caller.go",
        "pkg/caller/callret.go",
    ],
    importpath = "httpt/pkg/caller",
    deps = [
        ":pkg_services",
    ],
)

go_library(
    name = "pkg_config",
    srcs = [
        "pkg/config/config.go",
        "pkg/config/constants.go",
        "pkg/config/server_config.go",
    ],
    importpath = "httpt/pkg/config",
)

go_library(
    name = "pkg_logger",
    srcs = ["pkg/logger/logger.go"],
    importpath = "httpt/pkg/logger",
    deps = [
        "//deps/github.com/jeanphorn/log4go",
    ],
)

go_library(
    name = "pkg_router",
    srcs = ["pkg/router/router.go"],
    importpath = "httpt/pkg/router",
    deps = [
        ":pkg_caller",
        ":pkg_config",
        ":pkg_logger",
        ":pkg_uri",
        "//deps/github.com/valyala/fasthttp",
    ],
)

go_library(
    name = "pkg_server",
    srcs = ["pkg/server/server.go"],
    importpath = "httpt/pkg/server",
    deps = [
        ":pkg_config",
        ":pkg_logger",
        ":pkg_router",
        "//deps/github.com/valyala/fasthttp",
    ],
)

go_library(
    name = "pkg_services",
    srcs = [
        "pkg/services/manager.go",
        "pkg/services/services.go",
    ],
    importpath = "httpt/pkg/services",
)

go_library(
    name = "pkg_uri",
    srcs = [
        "pkg/uri/manager.go",
        "pkg/uri/uri.go",
    ],
    importpath = "httpt/pkg/uri",
    deps = [":pkg_services"],
)