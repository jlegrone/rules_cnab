load("@io_bazel_skydoc//stardoc:stardoc.bzl", "stardoc")
load("@bazel_skylib//:bzl_library.bzl", "bzl_library")

bzl_library(
    name = "docker-rules",
    srcs = ["@io_bazel_rules_docker//container:providers.bzl"],
)

stardoc(
    name = "docs",
    out = "README.md",
    format = "markdown",
    input = "cnab.bzl",
    deps = [":docker-rules"],
)
