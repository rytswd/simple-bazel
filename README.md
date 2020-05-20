# simple-bazel

This repository is mainly used for investigation (originally created for https://github.com/googleapis/google-cloud-go/issues/1898).

For the details as to which version of dependencies I used, please check out [WORKSPACE](./WORKSPACE).

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

## `issue/bazelbuild/bazel-gazelle/grpc-gateway-error` branch

Ref: https://github.com/bazelbuild/bazel-gazelle/issues/788

I simply added the following line

```go
import (
	...
	_ "github.com/grpc-ecosystem/grpc-gateway/runtime"
	...
)
```

This caused the following build error, at analysis phase:

```
$ bazel run simple-bazel
ERROR: (snip).../external/com_github_grpc_ecosystem_grpc_gateway/runtime/BUILD.bazel:5:1: no such target '@com_github_golang_protobuf//descriptor:go_default_library_gen': target 'go_default_library_gen' not declared in package 'descriptor' (did you mean 'go_default_library'?) defined by (snip).../external/com_github_golang_protobuf/descriptor/BUILD.bazel and referenced by '@com_github_grpc_ecosystem_grpc_gateway//runtime:go_default_library'
ERROR: (snip).../external/com_github_grpc_ecosystem_grpc_gateway/runtime/BUILD.bazel:5:1: no such target '@com_github_golang_protobuf//jsonpb:go_default_library_gen': target 'go_default_library_gen' not declared in package 'jsonpb' (did you mean 'go_default_library'?) defined by (snip).../external/com_github_golang_protobuf/jsonpb/BUILD.bazel and referenced by '@com_github_grpc_ecosystem_grpc_gateway//runtime:go_default_library'
ERROR: Analysis of target '//:simple-bazel' failed; build aborted: Analysis failed
INFO: Elapsed time: 31.438s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (177 packages loaded, 1368 targets configured)
FAILED: Build did NOT complete successfully (177 packages loaded, 1368 targets configured)
```

## `issue/google-cloud-go-1898` branch

**NOTE**: Fix has been merged to master.

<details>
<summary>Click to expand the details</summary>

Ref: https://github.com/googleapis/google-cloud-go/issues/1898

For overriding dependencies that are declared within [`go_rules_dependencies`](https://github.com/bazelbuild/rules_go/blob/master/go/private/repositories.bzl#L24), there is a [guideline on overriding dependencies at `bazelbuild/rules_go`](https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#overriding-dependencies).

However, even after following the steps mentioned to override `google.golang.org/genproto` version, I was getting the below error. This seems to indicate the version mismatch, and it may be configuration issue but could not see any other ways.

```shell
$ bazel run simple-bazel
INFO: Analyzed target //:simple-bazel (157 packages loaded, 7892 targets configured).
INFO: Found 1 target...
INFO: From Generating Descriptor Set proto_library @go_googleapis//google/spanner/v1:spanner_proto:
google/spanner/v1/keys.proto:21:1: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/mutation.proto:22:1: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/query_plan.proto:21:1: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/transaction.proto:22:1: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/type.proto:20:1: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/result_set.proto:24:1: warning: Import google/api/annotations.proto is unused.
INFO: From Generating into bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/spanner/v1/darwin_amd64_stripped/spanner_go_proto%/google.golang.org/genproto/googleapis/spanner/v1:
google/spanner/v1/keys.proto: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/mutation.proto: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/query_plan.proto: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/transaction.proto: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/type.proto: warning: Import google/api/annotations.proto is unused.
google/spanner/v1/result_set.proto: warning: Import google/api/annotations.proto is unused.
INFO: From Generating Descriptor Set proto_library @go_googleapis//google/iam/v1:iam_proto:
google/iam/v1/options.proto:20:1: warning: Import google/api/annotations.proto is unused.
google/iam/v1/policy.proto:21:1: warning: Import google/api/annotations.proto is unused.
INFO: From Generating into bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/iam/v1/darwin_amd64_stripped/iam_go_proto%/google.golang.org/genproto/googleapis/iam/v1:
google/iam/v1/options.proto: warning: Import google/api/annotations.proto is unused.
google/iam/v1/policy.proto: warning: Import google/api/annotations.proto is unused.
ERROR: /private/var/tmp/_bazel_ryota/d4ed1f7520bc209aa88b0b310e127cd6/external/com_google_cloud_go_spanner/BUILD.bazel:3:1: GoCompilePkg external/com_google_cloud_go_spanner/darwin_amd64_stripped/go_default_library%/cloud.google.com/go/spanner.a failed (Exit 1) builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder compilepkg -sdk external/go_sdk -installsuffix darwin_amd64 -src external/com_google_cloud_go_spanner/batch.go -src ... (remaining 103 argument(s) skipped)

Use --sandbox_debug to see verbose messages from the sandbox
compilepkg: error running subcommand: exit status 2
/private/var/tmp/_bazel_ryota/d4ed1f7520bc209aa88b0b310e127cd6/sandbox/darwin-sandbox/556/execroot/com_github_rytswd_simple_bazel/external/com_google_cloud_go_spanner/batch.go:185:3: unknown field 'QueryOptions' in struct literal of type "google.golang.org/genproto/googleapis/spanner/v1".ExecuteSqlRequest
/private/var/tmp/_bazel_ryota/d4ed1f7520bc209aa88b0b310e127cd6/sandbox/darwin-sandbox/556/execroot/com_github_rytswd_simple_bazel/external/com_google_cloud_go_spanner/client.go:286:20: undefined: "google.golang.org/genproto/googleapis/spanner/v1".ExecuteSqlRequest_QueryOptions
/private/var/tmp/_bazel_ryota/d4ed1f7520bc209aa88b0b310e127cd6/sandbox/darwin-sandbox/556/execroot/com_github_rytswd_simple_bazel/external/com_google_cloud_go_spanner/pdml.go:76:3: unknown field 'QueryOptions' in struct literal of type "google.golang.org/genproto/googleapis/spanner/v1".ExecuteSqlRequest
/private/var/tmp/_bazel_ryota/d4ed1f7520bc209aa88b0b310e127cd6/sandbox/darwin-sandbox/556/execroot/com_github_rytswd_simple_bazel/external/com_google_cloud_go_spanner/transaction.go:222:11: undefined: "google.golang.org/genproto/googleapis/spanner/v1".ExecuteSqlRequest_QueryOptions
/private/var/tmp/_bazel_ryota/d4ed1f7520bc209aa88b0b310e127cd6/sandbox/darwin-sandbox/556/execroot/com_github_rytswd_simple_bazel/external/com_google_cloud_go_spanner/transaction.go:230:13: undefined: "google.golang.org/genproto/googleapis/spanner/v1".ExecuteSqlRequest_QueryOptions
/private/var/tmp/_bazel_ryota/d4ed1f7520bc209aa88b0b310e127cd6/sandbox/darwin-sandbox/556/execroot/com_github_rytswd_simple_bazel/external/com_google_cloud_go_spanner/transaction.go:343:3: unknown field 'QueryOptions' in struct literal of type "google.golang.org/genproto/googleapis/spanner/v1".ExecuteSqlRequest
Target //:simple-bazel failed to build
Use --verbose_failures to see the command lines of failed build steps.
INFO: Elapsed time: 811.744s, Critical Path: 86.87s
INFO: 549 processes: 549 darwin-sandbox.
FAILED: Build did NOT complete successfully
FAILED: Build did NOT complete successfully
```

**Fix**:

- Compile `.proto` with protoc and commit (also correct some error)
- Run `go mod tidy` to pick up all dependencies
- Use `--build_file_proto_mode=disable_global` in `update-repos`

</details>
