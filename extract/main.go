package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/hmarr/pkgextract"
)

func main() {
	extractor := pkgextract.NewPackageExtractor(pkgextract.PackageExtractorOptions{
		InitialPkg: "cmd/go/internal/modload",
		OutputRoot: "./out",
		ImportRewriter: func(importPath string) string {
			newPath := strings.Replace(importPath, "internal", "_internal_", -1)
			return filepath.Join("github.com/dependabot/gomodules-extracted", newPath)
		},
		Filter: func(importPath string) bool {
			return strings.Contains(importPath, "internal")
		},
	})

	if err := extractor.Run(); err != nil {
		log.Fatal(err)
	}
}
