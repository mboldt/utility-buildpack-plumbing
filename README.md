# utility-buildpack-plumbing

A place to try out ideas for utility buildpacks build, package, test, release, libraries, frameworks, etc.

## Choices for Tooling

### `make` for Entry Point

`make` is the de facto standard across CNB projects
(
[lifecycle](https://github.com/buildpacks/lifecycle/blob/main/Makefile),
 [pack](https://github.com/buildpacks/pack/blob/main/Makefile),
 [libcnb](https://github.com/buildpacks/libcnb/blob/main/Makefile),
 [docs](https://github.com/buildpacks/docs/blob/main/Makefile),
 [samples](https://github.com/buildpacks/samples/blob/main/Makefile),
 etc.).

### [`pipeline-builder`](https://github.com/paketo-buildpacks/pipeline-builder) for CI/CD

`pipeline-builder` turns a small yaml file into all the common workflows for a buildpack.
It handles things like testing, preparing and drafting releases, publishing images, updating dependencies, etc.
`libcnb` uses `pipeline-builder` (see its [pipeline descriptor](https://github.com/buildpacks/libcnb/blob/main/.github/pipeline-descriptor.yml)).

### Go for Language

Again, Go is the de facto standard across CNB projects (lifecycle, pack, libcnb).
There are several Go libraries for developing buildpacks,
(e.g., [packit](https://github.com/paketo-buildpacks/packit),
 [libpak](https://github.com/paketo-buildpacks/libpak),
 [gcpbuildpack](https://github.com/GoogleCloudPlatform/buildpacks/tree/main/pkg/gcpbuildpack),
 [libbuildpack](https://github.com/cloudfoundry/libbuildpack)).

### `libcnb` for Library

To remain neutral towards third-party libraries, we will use `libcnb`.

### `sclevine/spec` for Test Framework

Another de facto standard across CNB projects (lifecycle, pack, libcnb).

## To Do / Open Questions

- Integration testing.
- Linting
- Use as a testbed for libcnbv2? Might have to deal with API changes.
  - Depends on which gets released first
  - Start with v1, use this as a migration guide?
- Monorepo utility buildpacks?
  - E.g., to use one binary to reduce the size of the buildpacks
  - Maybe not at first, since there will only be one buildpack; planting the seed for the future
