package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func tidyPath(pathComponents ...string) (string, error) {
	rawPath := filepath.Join(pathComponents...)
	currentUser, _ := user.Current()
	homeDirectory := currentUser.HomeDir
	if "~" == rawPath {
		return homeDirectory, nil
	} else if strings.HasPrefix(rawPath, "~/") {
		return filepath.Join(
				homeDirectory,
				strings.TrimPrefix(rawPath, "~/"),
			),
			nil
	}
	return filepath.Abs(rawPath)
}

var (
	pathTidier = tidyPath
)

func EnsureDirectoryExists(pathComponents ...string) error {
	combinedPath, err := pathTidier(pathComponents...)
	if nil == err {
		err = os.MkdirAll(combinedPath, os.ModePerm)
	}
	return err
}

func LoadFile(pathComponents ...string) ([]byte, error) {
	combinedPath, err := pathTidier(pathComponents...)
	if nil != err {
		return []byte{}, err
	}
	return ioutil.ReadFile(combinedPath)
}

func WriteFile(contents []byte, permissions os.FileMode, pathComponents ...string) error {
	combinedPath, err := pathTidier(pathComponents...)
	if nil != err {
		return err
	}
	return ioutil.WriteFile(combinedPath, contents, permissions)
}
