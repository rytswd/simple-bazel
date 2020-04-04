workspace(name = "com_github_rytswd_proj")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

### ---- http_archive ----------------------------------------------------- ###
# Go
# Ref: https://github.com/bazelbuild/rules_go#setup
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "142dd33e38b563605f0d20e89d9ef9eda0fc3cb539a14be1bdb1350de2eda659",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/v0.22.2/rules_go-v0.22.2.tar.gz",
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.22.2/rules_go-v0.22.2.tar.gz",
    ],
)

# Gazelle
# Ref: https://github.com/bazelbuild/bazel-gazelle#setup
http_archive(
    name = "bazel_gazelle",
    sha256 = "d8c45ee70ec39a57e7a05e5027c32b1576cc7f16d9dd37135b0eddde45cf1b10",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
    ],
)

# Protobuf
http_archive(
    name = "com_google_protobuf",
    strip_prefix = "protobuf-3.11.4",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.4.zip"],
    sha256 = "9748c0d90e54ea09e5e75fb7fac16edce15d2028d4356f32211cfa3c0e956564",
)
### ---- end of http_archive                                                ###
### ----------------------------------------------------------------------- ###

### ---- Load Go dependencies --------------------------------------------- ###
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# Override depenedency of `googel.golang.org/genproto`
# Ref: https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#overriding-dependencies
go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:MRHtG0U6SnaUb+s+LhNE1qt1FQ1wlhqr5E4usBKC0uA=",
    version = "v0.0.0-20200331122359-1ee6d9798940",
    # commit = "c50568487044f7393f957c36b91edbf731220cba",
)

# As per the above reference, this should be called after `go_repository` calls
go_rules_dependencies()

go_register_toolchains()
### ---- end of Go deps                                                     ###
### ----------------------------------------------------------------------- ###

### ---- Setup Gazelle ---------------------------------------------------- ###
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

# gazelle:repo bazel_gazelle

### ---- end of Gazelle setup                                               ###
### ----------------------------------------------------------------------- ###
