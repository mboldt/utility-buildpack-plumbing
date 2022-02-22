package utilitybuildpack

import (
	"fmt"

	"github.com/buildpacks/libcnb"
)

type Builder struct{}

func (b Builder) Build(context libcnb.BuildContext) (libcnb.BuildResult, error) {
	fmt.Println("--> Utility Buildpack")
	return libcnb.NewBuildResult(), nil
}
