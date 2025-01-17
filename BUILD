package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
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
    importpath = "httpt",
    deps = [
        ":pkg_config",
        ":pkg_logger",
        ":pkg_router",
        ":pkg_uri",
        "//deps/github.com/valyala/fasthttp",
    ],
)

go_library(
    name = "pkg_caller",
    srcs = [
        "pkg/caller/caller.go",
        "pkg/caller/response.go",
    ],
    importpath = "httpt/pkg/caller",
    deps = [
        ":pkg_config",
        ":pkg_runtime",
        ":pkg_status",
        ":pkg_uri",
        ":runtime_grpc",
        "//deps/github.com/satori/go.uuid",
        "@org_golang_google_grpc//:go_default_library",
        "//deps/github.com/valyala/fasthttp",
        # "@org_golang_google_grpc//codes:go_default_library",
        # "@org_golang_google_grpc//status:go_default_library",
    ],
)

go_library(
    name = "pkg_common",
    srcs = [
        "pkg/common/cocurrent_map.go",
        "pkg/common/type.go",
    ],
    importpath = "httpt/pkg/common",
)

go_library(
    name = "pkg_config",
    srcs = [
        "pkg/config/config.go",
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
    srcs = [
        "pkg/router/router.go",
        "pkg/router/json_router.go",
        "pkg/router/html_router.go",
    ],
    importpath = "httpt/pkg/router",
    deps = [
        ":pkg_caller",
        ":pkg_config",
        ":pkg_logger",
        ":pkg_uri",
        ":pkg_status",
        "//deps/github.com/valyala/fasthttp",
    ],
)

go_library(
    name = "pkg_runtime",
    srcs = [
        "pkg/runtime/manager.go",
        "pkg/runtime/runtimes.go",
    ],
    importpath = "httpt/pkg/runtime",
    deps = [
        ":pkg_common",
    ],
)

go_library(
    name = "pkg_status",
    srcs = [
        "pkg/status/status.go",
        "pkg/status/list.go",
    ],
    importpath = "httpt/pkg/status",
)

go_library(
    name = "pkg_uri",
    srcs = [
        "pkg/uri/manager.go",
        "pkg/uri/uri.go",
    ],
    importpath = "httpt/pkg/uri",
    deps = [
        ":pkg_common",
    ],
)

proto_library(
    name = "runtime_proto",
    srcs = ["build/protos/runtime.proto"],
)

go_proto_library(
    name = "runtime_grpc",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "httpt/build/protos/runtime",
    proto = ":runtime_proto",
)