Ref: https://github.com/bazelbuild/bazel-gazelle/issues/788

With the following patch setup:

```
go_repository(
    name = "com_github_golang_protobuf",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golang/protobuf",
    patch_args = ["-p1"],
    patches = ["@io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch"],
    sum = "h1:+Z5KGCizgyZCbGh1KZqA0fcLLkwbsjIzS4aV2v7wJX0=",
    version = "v1.4.2",
)
```

I got a lengthy error as below

```
$ bazel run simple-bazel
INFO: Repository com_github_golang_protobuf instantiated at:
  no stack (--record_rule_instantiation_callstack not enabled)
Repository rule go_repository defined at:
  /private/var/tmp/_bazel_Ryota/6984d376a01d0a19570e530ef54773df/external/bazel_gazelle/internal/go_repository.bzl:189:17: in <toplevel>
ERROR: An error occurred during the fetch of repository 'com_github_golang_protobuf':
   Traceback (most recent call last):
        File "/private/var/tmp/_bazel_Ryota/6984d376a01d0a19570e530ef54773df/external/bazel_gazelle/internal/go_repository.bzl", line 187
                patch(ctx)
        File "/private/var/tmp/_bazel_Ryota/6984d376a01d0a19570e530ef54773df/external/bazel_gazelle/internal/go_repository.bzl", line 277, in patch
                fail(<1 more arguments>)
Error applying patch @io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch:
patching file descriptor/BUILD.bazel
Hunk #1 succeeded at 14 (offset 3 lines).
patching file jsonpb/BUILD.bazel
Hunk #1 succeeded at 19 with fuzz 1 (offset 8 lines).
can't find file to patch at input line 47
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/jsonpb/jsonpb_test_proto/BUILD.bazel c/jsonpb/jsonpb_test_proto/BUILD.bazel
|--- b/jsonpb/jsonpb_test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/jsonpb/jsonpb_test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 65
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/proto/proto3_proto/BUILD.bazel c/proto/proto3_proto/BUILD.bazel
|--- b/proto/proto3_proto/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
|+++ c/proto/proto3_proto/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 80
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/proto/test_proto/BUILD.bazel c/proto/test_proto/BUILD.bazel
|--- b/proto/test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/proto/test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
patching file protoc-gen-go/descriptor/BUILD.bazel
patching file protoc-gen-go/generator/BUILD.bazel
Hunk #1 FAILED at 13.
1 out of 1 hunk FAILED -- saving rejects to file protoc-gen-go/generator/BUILD.bazel.rej
patching file protoc-gen-go/plugin/BUILD.bazel
can't find file to patch at input line 148
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/deprecated/BUILD.bazel c/protoc-gen-go/testdata/deprecated/BUILD.bazel
|--- b/protoc-gen-go/testdata/deprecated/BUILD.bazel    2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/deprecated/BUILD.bazel    2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 163
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_base/BUILD.bazel c/protoc-gen-go/testdata/extension_base/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_base/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_base/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 178
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_extra/BUILD.bazel c/protoc-gen-go/testdata/extension_extra/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_extra/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_extra/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 193
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_user/BUILD.bazel c/protoc-gen-go/testdata/extension_user/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_user/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_user/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 208
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/grpc/BUILD.bazel c/protoc-gen-go/testdata/grpc/BUILD.bazel
|--- b/protoc-gen-go/testdata/grpc/BUILD.bazel  2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/grpc/BUILD.bazel  2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 223
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/import_public/BUILD.bazel c/protoc-gen-go/testdata/import_public/BUILD.bazel
|--- b/protoc-gen-go/testdata/import_public/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/import_public/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 241
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/import_public/sub/BUILD.bazel c/protoc-gen-go/testdata/import_public/sub/BUILD.bazel
|--- b/protoc-gen-go/testdata/import_public/sub/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/import_public/sub/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 259
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/BUILD.bazel c/protoc-gen-go/testdata/imports/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 278
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/fmt/BUILD.bazel c/protoc-gen-go/testdata/imports/fmt/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/fmt/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/fmt/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 293
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel c/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 311
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel c/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 329
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel c/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 347
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/multi/BUILD.bazel c/protoc-gen-go/testdata/multi/BUILD.bazel
|--- b/protoc-gen-go/testdata/multi/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/multi/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 366
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/my_test/BUILD.bazel c/protoc-gen-go/testdata/my_test/BUILD.bazel
|--- b/protoc-gen-go/testdata/my_test/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/my_test/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 381
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/proto3/BUILD.bazel c/protoc-gen-go/testdata/proto3/BUILD.bazel
|--- b/protoc-gen-go/testdata/proto3/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/proto3/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
patching file ptypes/BUILD.bazel
Hunk #1 succeeded at 20 (offset 2 lines).
patching file ptypes/any/BUILD.bazel
patching file ptypes/duration/BUILD.bazel
patching file ptypes/empty/BUILD.bazel
patching file ptypes/struct/BUILD.bazel
patching file ptypes/timestamp/BUILD.bazel
patching file ptypes/wrappers/BUILD.bazel
ERROR: /private/var/tmp/_bazel_Ryota/6984d376a01d0a19570e530ef54773df/external/com_google_cloud_go_spanner/BUILD.bazel:3:1: @com_google_cloud_go_spanner//:go_default_library depends on @com_github_golang_protobuf//ptypes/struct:go_default_library in repository @com_github_golang_protobuf which failed to fetch. no such package '@com_github_golang_protobuf//ptypes/struct': Traceback (most recent call last):
        File "/private/var/tmp/_bazel_Ryota/6984d376a01d0a19570e530ef54773df/external/bazel_gazelle/internal/go_repository.bzl", line 187
                patch(ctx)
        File "/private/var/tmp/_bazel_Ryota/6984d376a01d0a19570e530ef54773df/external/bazel_gazelle/internal/go_repository.bzl", line 277, in patch
                fail(<1 more arguments>)
Error applying patch @io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch:
patching file descriptor/BUILD.bazel
Hunk #1 succeeded at 14 (offset 3 lines).
patching file jsonpb/BUILD.bazel
Hunk #1 succeeded at 19 with fuzz 1 (offset 8 lines).
can't find file to patch at input line 47
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/jsonpb/jsonpb_test_proto/BUILD.bazel c/jsonpb/jsonpb_test_proto/BUILD.bazel
|--- b/jsonpb/jsonpb_test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/jsonpb/jsonpb_test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 65
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/proto/proto3_proto/BUILD.bazel c/proto/proto3_proto/BUILD.bazel
|--- b/proto/proto3_proto/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
|+++ c/proto/proto3_proto/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 80
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/proto/test_proto/BUILD.bazel c/proto/test_proto/BUILD.bazel
|--- b/proto/test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/proto/test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
patching file protoc-gen-go/descriptor/BUILD.bazel
patching file protoc-gen-go/generator/BUILD.bazel
Hunk #1 FAILED at 13.
1 out of 1 hunk FAILED -- saving rejects to file protoc-gen-go/generator/BUILD.bazel.rej
patching file protoc-gen-go/plugin/BUILD.bazel
can't find file to patch at input line 148
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/deprecated/BUILD.bazel c/protoc-gen-go/testdata/deprecated/BUILD.bazel
|--- b/protoc-gen-go/testdata/deprecated/BUILD.bazel    2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/deprecated/BUILD.bazel    2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 163
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_base/BUILD.bazel c/protoc-gen-go/testdata/extension_base/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_base/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_base/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 178
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_extra/BUILD.bazel c/protoc-gen-go/testdata/extension_extra/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_extra/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_extra/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 193
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_user/BUILD.bazel c/protoc-gen-go/testdata/extension_user/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_user/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_user/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 208
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/grpc/BUILD.bazel c/protoc-gen-go/testdata/grpc/BUILD.bazel
|--- b/protoc-gen-go/testdata/grpc/BUILD.bazel  2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/grpc/BUILD.bazel  2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 223
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/import_public/BUILD.bazel c/protoc-gen-go/testdata/import_public/BUILD.bazel
|--- b/protoc-gen-go/testdata/import_public/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/import_public/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 241
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/import_public/sub/BUILD.bazel c/protoc-gen-go/testdata/import_public/sub/BUILD.bazel
|--- b/protoc-gen-go/testdata/import_public/sub/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/import_public/sub/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 259
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/BUILD.bazel c/protoc-gen-go/testdata/imports/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 278
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/fmt/BUILD.bazel c/protoc-gen-go/testdata/imports/fmt/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/fmt/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/fmt/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 293
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel c/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 311
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel c/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 329
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel c/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 347
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/multi/BUILD.bazel c/protoc-gen-go/testdata/multi/BUILD.bazel
|--- b/protoc-gen-go/testdata/multi/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/multi/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 366
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/my_test/BUILD.bazel c/protoc-gen-go/testdata/my_test/BUILD.bazel
|--- b/protoc-gen-go/testdata/my_test/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/my_test/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 381
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/proto3/BUILD.bazel c/protoc-gen-go/testdata/proto3/BUILD.bazel
|--- b/protoc-gen-go/testdata/proto3/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/proto3/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
patching file ptypes/BUILD.bazel
Hunk #1 succeeded at 20 (offset 2 lines).
patching file ptypes/any/BUILD.bazel
patching file ptypes/duration/BUILD.bazel
patching file ptypes/empty/BUILD.bazel
patching file ptypes/struct/BUILD.bazel
patching file ptypes/timestamp/BUILD.bazel
patching file ptypes/wrappers/BUILD.bazel
ERROR: Analysis of target '//:simple-bazel' failed; build aborted: no such package '@com_github_golang_protobuf//ptypes/struct': Traceback (most recent call last):
        File "/private/var/tmp/_bazel_Ryota/6984d376a01d0a19570e530ef54773df/external/bazel_gazelle/internal/go_repository.bzl", line 187
                patch(ctx)
        File "/private/var/tmp/_bazel_Ryota/6984d376a01d0a19570e530ef54773df/external/bazel_gazelle/internal/go_repository.bzl", line 277, in patch
                fail(<1 more arguments>)
Error applying patch @io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch:
patching file descriptor/BUILD.bazel
Hunk #1 succeeded at 14 (offset 3 lines).
patching file jsonpb/BUILD.bazel
Hunk #1 succeeded at 19 with fuzz 1 (offset 8 lines).
can't find file to patch at input line 47
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/jsonpb/jsonpb_test_proto/BUILD.bazel c/jsonpb/jsonpb_test_proto/BUILD.bazel
|--- b/jsonpb/jsonpb_test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/jsonpb/jsonpb_test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 65
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/proto/proto3_proto/BUILD.bazel c/proto/proto3_proto/BUILD.bazel
|--- b/proto/proto3_proto/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
|+++ c/proto/proto3_proto/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 80
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/proto/test_proto/BUILD.bazel c/proto/test_proto/BUILD.bazel
|--- b/proto/test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/proto/test_proto/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
patching file protoc-gen-go/descriptor/BUILD.bazel
patching file protoc-gen-go/generator/BUILD.bazel
Hunk #1 FAILED at 13.
1 out of 1 hunk FAILED -- saving rejects to file protoc-gen-go/generator/BUILD.bazel.rej
patching file protoc-gen-go/plugin/BUILD.bazel
can't find file to patch at input line 148
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/deprecated/BUILD.bazel c/protoc-gen-go/testdata/deprecated/BUILD.bazel
|--- b/protoc-gen-go/testdata/deprecated/BUILD.bazel    2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/deprecated/BUILD.bazel    2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 163
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_base/BUILD.bazel c/protoc-gen-go/testdata/extension_base/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_base/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_base/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 178
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_extra/BUILD.bazel c/protoc-gen-go/testdata/extension_extra/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_extra/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_extra/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 193
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/extension_user/BUILD.bazel c/protoc-gen-go/testdata/extension_user/BUILD.bazel
|--- b/protoc-gen-go/testdata/extension_user/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/extension_user/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 208
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/grpc/BUILD.bazel c/protoc-gen-go/testdata/grpc/BUILD.bazel
|--- b/protoc-gen-go/testdata/grpc/BUILD.bazel  2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/grpc/BUILD.bazel  2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 223
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/import_public/BUILD.bazel c/protoc-gen-go/testdata/import_public/BUILD.bazel
|--- b/protoc-gen-go/testdata/import_public/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/import_public/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 241
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/import_public/sub/BUILD.bazel c/protoc-gen-go/testdata/import_public/sub/BUILD.bazel
|--- b/protoc-gen-go/testdata/import_public/sub/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/import_public/sub/BUILD.bazel     2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 259
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/BUILD.bazel c/protoc-gen-go/testdata/imports/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 278
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/fmt/BUILD.bazel c/protoc-gen-go/testdata/imports/fmt/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/fmt/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/fmt/BUILD.bazel   2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 293
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel c/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_a_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 311
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel c/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_a_2/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 329
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel c/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel
|--- b/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/imports/test_b_1/BUILD.bazel      2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 347
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/multi/BUILD.bazel c/protoc-gen-go/testdata/multi/BUILD.bazel
|--- b/protoc-gen-go/testdata/multi/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/multi/BUILD.bazel 2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 366
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/my_test/BUILD.bazel c/protoc-gen-go/testdata/my_test/BUILD.bazel
|--- b/protoc-gen-go/testdata/my_test/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/my_test/BUILD.bazel       2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
can't find file to patch at input line 381
Perhaps you used the wrong -p or --strip option?
The text leading up to this was:
--------------------------
|diff -urN b/protoc-gen-go/testdata/proto3/BUILD.bazel c/protoc-gen-go/testdata/proto3/BUILD.bazel
|--- b/protoc-gen-go/testdata/proto3/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
|+++ c/protoc-gen-go/testdata/proto3/BUILD.bazel        2000-01-01 00:00:00.000000000 -0000
--------------------------
File to patch:
Skip this patch? [y]
Skipping patch.
1 out of 1 hunk ignored
patching file ptypes/BUILD.bazel
Hunk #1 succeeded at 20 (offset 2 lines).
patching file ptypes/any/BUILD.bazel
patching file ptypes/duration/BUILD.bazel
patching file ptypes/empty/BUILD.bazel
patching file ptypes/struct/BUILD.bazel
patching file ptypes/timestamp/BUILD.bazel
patching file ptypes/wrappers/BUILD.bazel
INFO: Elapsed time: 1.476s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (19 packages loaded, 3 targets configured)
FAILED: Build did NOT complete successfully (19 packages loaded, 3 targets configured)
    Fetching @local_config_xcode; fetching
```
