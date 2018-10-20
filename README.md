# Go modules internals

The internals of Go modules, extracted from [golang/go](https://github.com/golang/go). All
occurrences of `internal` in import paths have been rewritten to `_internal_`,
but the code is otherwise unchanged.

## Re-extracting

The `extract/` directory contains the code to do this. Run the Go program, then
copy the contents of the `out/github.com/dependabot/gomodules-extracted`
directory to the top level of this repo.
