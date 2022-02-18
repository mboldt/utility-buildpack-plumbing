.PHONY: test
test: build
	pack build --verbose --path testdata/app --builder cnbs/sample-builder:alpine --buildpack ./out sample-utility-buidpack-app

.PHONY: build
build: out out/buildpack.toml out/bin/build out/bin/detect

out:
	mkdir out

out/buildpack.toml: buildpack.toml
	cp buildpack.toml out/buildpack.toml

out/bin/build: cmd/build/main.go
	GOOS=linux GOARCH=amd64 go build -o out/bin/build cmd/build/main.go

out/bin/detect: cmd/detect/main.go
	GOOS=linux GOARCH=amd64 go build -o out/bin/detect cmd/detect/main.go

.PHONY: clean
clean:
	rm -rf out/
