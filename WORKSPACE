workspace(name = "httpt")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "08c3cd71857d58af3cda759112437d9e63339ac9c6e0042add43f4d94caf632d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.2/rules_go-v0.24.2.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.2/rules_go-v0.24.2.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains(
    go_version = "1.15.2",
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "d4113967ab451dd4d2d767c3ca5f927fec4b30f3b2c6f8135a2033b9c05a5687",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
    ],
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

# gRPC v1.27.3
http_archive(
    name = "com_github_grpc_grpc",
    urls = [
        "https://github.com/grpc/grpc/archive/e73882dc0fcedab1ffe789e44ed6254819639ce3.tar.gz",
    ],
    strip_prefix = "grpc-e73882dc0fcedab1ffe789e44ed6254819639ce3",
)
load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")
grpc_deps()

load("//:build/workspace.bzl", "httpt_dependencies")
httpt_dependencies()