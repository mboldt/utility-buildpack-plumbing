package utilitybuildpack

import (
	"fmt"

	"github.com/buildpacks/libcnb"
)

type Detector struct{}

func (d Detector) Detect(context libcnb.DetectContext) (libcnb.DetectResult, error) {
	fmt.Println("--> Utility Buildpack")
	return libcnb.DetectResult{Pass: true}, nil
}
