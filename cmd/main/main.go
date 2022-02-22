package main

import (
	"utility-buildpack-plumbing/utilitybuildpack"

	"github.com/buildpacks/libcnb"
)

func main() {
	libcnb.Main(
		utilitybuildpack.Detector{},
		utilitybuildpack.Builder{},
	)
}
