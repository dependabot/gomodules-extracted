# Go modules internals

The internals of Go modules, extracted from [golang/go](https://github.com/golang/go). All
occurrences of `internal` in import paths have been rewritten to `_internal_`, which makes
these packages importable, but everything is otherwise unchanged.

## Re-extracting

Run `script/extract`.
