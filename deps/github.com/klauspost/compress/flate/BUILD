load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "flate",
    srcs = [
        "deflate.go",
        "dict_decoder.go",
        "fast_encoder.go",
        "huffman_bit_writer.go",
        "huffman_code.go",
        "huffman_sortByFreq.go",
        "huffman_sortByLiteral.go",
        "inflate.go",
        "inflate_gen.go",
        "level1.go",
        "level2.go",
        "level3.go",
        "level4.go",
        "level5.go",
        "level6.go",
        "stateless.go",
        "token.go",
    ],
    importpath = "github.com/klauspost/compress/flate",
    visibility = ["//visibility:public"],
)

go_test(
    name = "flate_test",
    srcs = [
        "deflate_test.go",
        "dict_decoder_test.go",
        "flate_test.go",
        "huffman_bit_writer_test.go",
        "inflate_test.go",
        "reader_test.go",
        "token_test.go",
        "writer_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":flate"],
)
