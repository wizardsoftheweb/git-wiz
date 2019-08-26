OUT := git-wiz
PKG := github.com/wizardsoftheweb/git-wiz
VERSION := $(shell git describe --always --long --dirty)

clean:
	@rm -rf ./build

build-version:
	@mkdir -p build
	go build -i -v -o build/${OUT} -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

build: clean build-version
