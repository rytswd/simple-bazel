# simple-bazel

This repository is mainly used for investigation (originally created for https://github.com/googleapis/google-cloud-go/issues/1898).

For the details as to which version of dependencies I used, please check out [WORKSPACE](./blob/master/WORKSPACE).

## `master` branch

In the `master` branch, it is confirmed that Bazel setup works, as long as no override is needed from the underlying code.

```shell
$ bazel run simple-bazel
INFO: Analyzed target //:simple-bazel (112 packages loaded, 1082 targets configured).
INFO: Found 1 target...
Target //:simple-bazel up-to-date:
  bazel-bin/darwin_amd64_stripped/simple-bazel
INFO: Elapsed time: 872.241s, Critical Path: 81.19s
INFO: 516 processes: 516 darwin-sandbox.
INFO: Build completed successfully, 533 total actions
INFO: Build completed successfully, 533 total actions
hello
```

## `issue/google-cloud-go-1898` branch

Ref: https://github.com/googleapis/google-cloud-go/issues/1898

For overriding dependencies that are declared within [`go_rules_dependencies`](https://github.com/bazelbuild/rules_go/blob/master/go/private/repositories.bzl#L24), there is a guideline on [Overriding dependencies in `bazelbuild/rules_go`](https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#overriding-dependencies).

However, even after following the steps mentioned to override `google.golang.org/genproto` version, I was getting the same error.
