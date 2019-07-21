load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/goharbor/harbor/cnab/bundlegen",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "bundlegen",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
