BINARY_NAME := git-wiz
PKG := github.com/wizardsoftheweb/git-wiz
SEMVER := $(shell git describe --abbrev=0)
VERSION := $(shell git describe --always --long --dirty)

windows_386:
	env GOOS=windows GOARCH=386 go build -i -v -o build/${BINARY_NAME}_${SEMVER}_windows_386 -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

windows_amd64:
	env GOOS=windows GOARCH=amd64 go build -i -v -o build/${BINARY_NAME}_${SEMVER}_windows_amd64 -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

linux_386:
	env GOOS=linux GOARCH=386 go build -i -v -o build/${BINARY_NAME}_${SEMVER}_linux_386 -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -i -v -o build/${BINARY_NAME}_${SEMVER}_linux_amd64 -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

darwin_amd64:
	env GOOS=darwin GOARCH=amd64 go build -i -v -o build/${BINARY_NAME}_${SEMVER}_darwin_amd64 -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

clean:
	@rm -rf ./build

build-version:
	go build -i -v -o "build/${BINARY_NAME}" -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

build: clean build-version


crossbuild: clean windows_386 windows_amd64 linux_386 linux_amd64 darwin_amd64

