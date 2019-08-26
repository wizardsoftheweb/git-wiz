BINARY_NAME := git-wiz
PKG := github.com/wizardsoftheweb/git-wiz
SEMVER := $(shell git describe --abbrev=0)
VERSION := $(shell git describe --always --long --dirty)

clean:
	@rm -rf ./build

build-version:
	go build -i -v -o "build/${BINARY_NAME}" -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

build: clean build-version
