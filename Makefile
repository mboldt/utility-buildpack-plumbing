.PHONY: test
test:
	pack build --path testdata/app --builder cnbs/sample-builder:alpine --buildpack ./buildpack sample-utility-buidpack-app
