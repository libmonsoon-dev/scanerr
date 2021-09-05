package packages

import "golang.org/x/tools/go/packages"

func Count(pkgs []*packages.Package) (n int) {
	packages.Visit(pkgs, nil, func(_ *packages.Package) {
		n++
	})
	return
}
