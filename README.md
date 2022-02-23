# utility-buildpack-plumbing

A place to try out ideas for utility buildpacks build, package, test, release, libraries, frameworks, etc.

## Choices for Tooling

## `make` for Entry Point

`make` is the de facto standard across CNB projects
(
[lifecycle](https://github.com/buildpacks/lifecycle/blob/main/Makefile),
 [pack](https://github.com/buildpacks/pack/blob/main/Makefile),
 [libcnb](https://github.com/buildpacks/libcnb/blob/main/Makefile),
 [docs](https://github.com/buildpacks/docs/blob/main/Makefile),
 [samples](https://github.com/buildpacks/samples/blob/main/Makefile),
 etc.).

## GitHub Actions for CI/CD

Again, GitHub Actions is the de facto standard across CNB projects
([lifecycle](https://github.com/buildpacks/lifecycle/tree/main/.github/workflows),
 [pack](https://github.com/buildpacks/pack/tree/main/.github/workflows),
 [libcnb](https://github.com/buildpacks/libcnb/tree/main/.github/workflows),
 [docs](https://github.com/buildpacks/docs/tree/main/.github/workflows),
 [samples](https://github.com/buildpacks/samples/tree/main/.github/workflows),
 etc.).


## Go for Language

Again, Go is the de facto standard across CNB projects (lifecycle, pack, libcnb).
There are several Go libraries for developing buildpacks,
(e.g., [packit](https://github.com/paketo-buildpacks/packit),
 [libpak](https://github.com/paketo-buildpacks/libpak),
 [gcpbuildpack](https://github.com/GoogleCloudPlatform/buildpacks/tree/main/pkg/gcpbuildpack),
 [libbuildpack](https://github.com/cloudfoundry/libbuildpack)).

## `libcnb` for Library

To remain unopinionated towards third-party libraries, we will use `libcnb`

## `sclevine/spec` for Test Framework

Another de facto standard across CNB projects (lifecycle, pack, libcnb).
