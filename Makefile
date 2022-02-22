.PHONY: test
test: package
	pack build --verbose --path testdata/app --builder cnbs/sample-builder:alpine --buildpack utility-buildpack sample-utility-buidpack-app

.PHONY: build
build: out out/buildpack.toml out/bin/main out/bin/build out/bin/detect

out:
	mkdir out

out/buildpack.toml: buildpack.toml
	cp buildpack.toml out/buildpack.toml

out/bin/main: cmd/main/main.go
	GOOS=linux GOARCH=amd64 go build -o out/bin/main cmd/main/main.go

out/bin/build: out/bin/main
	ln -sf main out/bin/build

out/bin/detect: out/bin/main
	ln -sf main out/bin/detect

.PHONY: package
package: build
	pack buildpack package utility-buildpack --config ./package.toml

.PHONY: clean
clean:
	rm -rf out/
