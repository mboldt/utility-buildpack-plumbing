package utilitybuildpack_test

import (
	"path/filepath"
	"testing"
	"utility-buildpack-plumbing/utilitybuildpack"

	"github.com/buildpacks/libcnb"
	"github.com/sclevine/spec"
)

func testDetect(t *testing.T, when spec.G, it spec.S) {
	var (
		context  libcnb.DetectContext      = libcnb.DetectContext{}
		detector utilitybuildpack.Detector = utilitybuildpack.Detector{}
	)

	when("there is no .fail-detect file in the app dir", func() {
		it.Before(func() {
			context.Application = libcnb.Application{
				Path: filepath.Join("..", "testdata", "app"),
			}
		})

		it("passes detection", func() {
			result, err := detector.Detect(context)
			if err != nil {
				t.Error("Unexpected error", err)
			}
			if !result.Pass {
				t.Errorf("Expected %s to pass detection", context.Application.Path)
			}
		})
	})

	when("there is a .fail-detect file in the app dir", func() {
		it.Before(func() {
			context.Application = libcnb.Application{
				Path: filepath.Join("..", "testdata", "app-fail-detect"),
			}
		})

		it("fails detection", func() {
			result, err := detector.Detect(context)
			if err != nil {
				t.Error("Unexpected error:", err)
			}
			if result.Pass {
				t.Errorf("Expected %s to fail detection", context.Application.Path)
			}
		})
	})
}
