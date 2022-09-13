workspace(name = "rules_go_simple")

load("@rules_go_simple//:deps.bzl", "go_rules_dependencies", "hello_repo")


hello_repo(
    name = "hello",
    message = "Hello, world!",
)

go_rules_dependencies()