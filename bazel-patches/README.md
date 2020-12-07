### grpc-gateway.patch

For generating the patch file, run the following command:

```bash
{
    # Create a directory to prepare a patch
    mkdir /tmp/grpc-gateway-patch/
    cd /tmp/grpc-gateway-patch/

    # Shallow clone based on version
    git clone --depth 1 --branch v2.0.1 https://github.com/grpc-ecosystem/grpc-gateway.git
    cd ./grpc-gateway

    # (Optional)
    # Remove bazel version if you need to run using different Bazel version
    rm .bazelversion

    # Add Gazelle directive to disable proto compilation
    echo '# gazelle:proto disable_global' >> BUILD

    # Run Gazelle update-repos command to update repositories.bzl with
    # disable_global flag
    bazel run gazelle -- update-repos \
        -from_file=go.mod \
        -to_macro=repositories.bzl%go_repositories \
        --build_file_proto_mode=disable_global

    # Remove BUILD.bazel file with conflicting import
    rm runtime/BUILD.bazel
    rm runtime/internal/examplepb/BUILD.bazel

    # Run Gazelle fix command to regenerate BUILD.bazel based on diasble_global
    bazel run gazelle -- fix

    # Create a patch file for two files that causes the build error:
    #   - `repositories.bzl`
    #   - `runtime/internal/examplepb/BUILD.bazel`
    git diff -u repositories.bzl runtime/BUILD.bazel runtime/internal/examplepb/BUILD.bazel > ../grpc-gateway.patch

    # Get back and clean up shallow clone
    cd ../
    rm -rf ./grpc-gateway
}
```
