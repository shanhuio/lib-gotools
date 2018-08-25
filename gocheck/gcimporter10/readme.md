Ported from `go/internal/gcimporter` in Golang source code 1.10.

This might break in the future. To maintain this, need to track all
the future changes in the `gcimporter` package.

We ported this to add a `build.Context` into the importer, so that
it does not always use the default context.
