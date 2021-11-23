# Go modules internals

The internals of Go modules, extracted from [golang/go](https://github.com/golang/go). All
occurrences of `internal` in import paths have been rewritten to `_internal_`, which makes
these packages importable, but everything is otherwise unchanged.

This code was originally used in `dependabot-core`, but that was removed by https://github.com/dependabot/dependabot-core/pull/4434/.

## Re-extracting

Run `script/extract`.
