load("//internal:rules.bzl",
 _go_binary = "go_binary", 
 _go_library = "go_library")

load(
    "//internal:providers.bzl",
    _GoLibraryInfo = "GoLibraryInfo"
)
go_binary = _go_binary
go_library = _go_library
GoLibraryInfo = _GoLibraryInfo