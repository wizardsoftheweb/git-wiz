OUT := git-wiz
PKG := github.com/wizardsoftheweb/git-wiz
VERSION := $(shell git describe --always --long --dirty)

clean:
	@rm -rf ./${OUT}

build-version:
	go build -i -v -o ${OUT} -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

build: clean build-version
