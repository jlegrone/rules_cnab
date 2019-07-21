load(
    "@io_bazel_rules_docker//container:providers.bzl",
    "PushInfo",
)

def _push_info_to_json(p):
    return struct(
        registry = p.registry,
        repository = p.repository,
        tag = p.tag,
        digestPath = p.digest.path,
    ).to_json()

def _cnab_impl(ctx):
    args = ["-bundle-path", ctx.outputs.bundle.path]
    deps = []

    invocation_images_arg = "-invocation-images="
    for _, target in enumerate(ctx.attr.invocation_images):
        deps.append(target[PushInfo].digest)
        invocation_images_arg += _push_info_to_json(target[PushInfo]) + "\n"
    args.append(invocation_images_arg)

    images_arg = "-images="
    for target, name in ctx.attr.images.items():
        images_arg += "{}={}\n".format(name, _push_info_to_json(target[PushInfo]))
        deps.append(target[PushInfo].digest)
    args.append(images_arg)

    ctx.actions.run(
        inputs = deps,
        outputs = [ctx.outputs.bundle],
        arguments = args,
        progress_message = "Generating bundle.json file: " + ctx.label.name,
        executable = ctx.executable._bundle_writer,
    )

cnab = rule(
    implementation = _cnab_impl,
    attrs = {
        "invocation_images": attr.label_list(
            mandatory = True,
            providers = [PushInfo],
            doc = "The array of invocation image definitions for this bundle.",
        ),
        "images": attr.label_keyed_string_dict(
            allow_empty = True,
            providers = [PushInfo],
            doc = "Images that are used by this bundle.",
        ),
        "_bundle_writer": attr.label(
            executable = True,
            cfg = "host",
            allow_files = True,
            default = Label("//bundlegen:bundlegen"),
        ),
    },
    outputs = {"bundle": "%{name}.bundle.json"},
)
