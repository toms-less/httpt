load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "zstd",
    srcs = [
        "bitreader.go",
        "bitwriter.go",
        "blockdec.go",
        "blockenc.go",
        "blocktype_string.go",
        "bytebuf.go",
        "bytereader.go",
        "decoder.go",
        "decoder_options.go",
        "dict.go",
        "enc_better.go",
        "enc_dfast.go",
        "enc_fast.go",
        "enc_params.go",
        "encoder.go",
        "encoder_options.go",
        "framedec.go",
        "frameenc.go",
        "fse_decoder.go",
        "fse_encoder.go",
        "fse_predefined.go",
        "hash.go",
        "history.go",
        "seqdec.go",
        "seqenc.go",
        "snappy.go",
        "zstd.go",
    ],
    importpath = "github.com/klauspost/compress/zstd",
    visibility = ["//visibility:public"],
    deps = [
        "//deps/github.com/klauspost/compress/huff0",
        "//deps/github.com/klauspost/compress/snappy",
        "//deps/github.com/klauspost/compress/zstd/internal/xxhash",
    ],
)

go_test(
    name = "zstd_test",
    srcs = [
        "decoder_test.go",
        "dict_test.go",
        "encoder_options_test.go",
        "encoder_test.go",
        "snappy_test.go",
        "zstd_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":zstd"],
    deps = [
        "//deps/github.com/klauspost/compress/snappy",
        "//deps/github.com/klauspost/compress/zip",
        "//deps/github.com/klauspost/compress/zstd/internal/xxhash",
    ],
)
