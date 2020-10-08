// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modfetch

import (
	"github.com/dependabot/gomodules-extracted/cmd/go/_internal_/cfg"
	"github.com/dependabot/gomodules-extracted/cmd/go/_internal_/get"
	"github.com/dependabot/gomodules-extracted/cmd/go/_internal_/str"
)

// allowInsecure reports whether we are allowed to fetch this path in an insecure manner.
func allowInsecure(path string) bool {
	return get.Insecure || str.GlobsMatchPath(cfg.GOINSECURE, path)
}
