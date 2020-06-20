package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")
load("//:build/workspace.bzl", "RELEASE_VERSION")

#build package of dashboard.
pkg_tar(
    name = "httpt-" + RELEASE_VERSION,
    extension = "tar.gz",
    srcs = [":httpt"],
    mode = "0755",
)

go_binary(
    name = "httpt",
    embed = [":main"],
)

go_library(
    name = "main",
    srcs = [
        "app.go"
    ],
    importpath = "httpt",
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
        ":pkg_services"
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
    srcs = [
        "pkg/logger/logger.go",
    ],
    importpath = "httpt/pkg/logger",
    deps = [
        ":vendor_log4go"
    ],
)

go_library(
    name = "pkg_router",
    srcs = [
        "pkg/router/router.go",
    ],
    importpath = "httpt/pkg/router",
    deps = [
        ":pkg_caller",
        ":pkg_config",
        ":pkg_logger",
        ":pkg_uri",
        "@com_github_valyala_fasthttp//:go_default_library",
    ],
)

go_library(
    name = "pkg_server",
    srcs = [
        "pkg/server/server.go",
    ],
    importpath = "httpt/pkg/server",
    deps = [
        ":pkg_config",
        ":pkg_logger",
        ":pkg_router",
        "@com_github_valyala_fasthttp//:go_default_library",
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

go_library(
    name = "vendor_log4go",
    srcs = [
        "vendor/github.com/jeanphorn/log4go/category.go",
        "vendor/github.com/jeanphorn/log4go/filelog.go",
        "vendor/github.com/jeanphorn/log4go/jsonconfig.go",
        "vendor/github.com/jeanphorn/log4go/log4go.go",
        "vendor/github.com/jeanphorn/log4go/pattlog.go",
        "vendor/github.com/jeanphorn/log4go/socklog.go",
        "vendor/github.com/jeanphorn/log4go/termlog.go",
        "vendor/github.com/jeanphorn/log4go/util.go",
        "vendor/github.com/jeanphorn/log4go/wrapper.go",
        "vendor/github.com/jeanphorn/log4go/xmlconfig.go",
    ],
    importpath = "github.com/jeanphorn/log4go",
    deps = [":vendor_tookits_file"],
)

go_library(
    name = "vendor_tookits_file",
    srcs = [
        "vendor/github.com/toolkits/file/downloader.go",
        "vendor/github.com/toolkits/file/file.go",
        "vendor/github.com/toolkits/file/reader.go",
        "vendor/github.com/toolkits/file/writer.go",
    ],
    importpath = "github.com/toolkits/file",
)