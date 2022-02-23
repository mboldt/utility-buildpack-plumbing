package utilitybuildpack_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnit(t *testing.T) {
	suite := spec.New("Utility Buildpack", spec.Report(report.Terminal{}))
	suite("Detect", testDetect)
	suite.Run(t)
}
