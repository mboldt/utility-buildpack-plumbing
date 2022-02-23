package utilitybuildpack

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/buildpacks/libcnb"
	"github.com/pkg/errors"
)

type Detector struct{}

func (d Detector) Detect(context libcnb.DetectContext) (libcnb.DetectResult, error) {
	fmt.Println("--> Utility Buildpack")
	_, err := os.Stat(filepath.Join(context.Application.Path, ".fail-detect"))
	if err == nil {
		return libcnb.DetectResult{Pass: false}, nil
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return libcnb.DetectResult{Pass: false}, errors.Wrap(err, "Failed to stat .fail-detect")
	}
	return libcnb.DetectResult{Pass: true}, nil
}
