load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "xxhash",
    srcs = [
        "xxhash.go",
        "xxhash_amd64.go",
        "xxhash_amd64.s",
        "xxhash_other.go",
        "xxhash_safe.go",
    ],
    importpath = "github.com/klauspost/compress/zstd/internal/xxhash",
    visibility = ["//deps/github.com/klauspost/compress/zstd:__subpackages__"],
)

go_test(
    name = "xxhash_test",
    srcs = ["xxhash_test.go"],
    embed = [":xxhash"],
)
