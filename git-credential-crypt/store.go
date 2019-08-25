package main

import (
	"os"
	"strings"
)

var DefaultStoreFileLocations = [][]string{
	{"~", ".git-credentials"},
	{os.Getenv("XDG_CONFIG_HOME"), "git", "credentials"},
	{"~", ".config", "git", "credentials"},
}

type Store struct {
	FileName       string
	Sites          []*Site
	availableFiles []string
}

func NewStoreFromDisk(filename string) *Store {
	cleanedFilename, _ := tidyPath(filename)
	store := Store{
		FileName: cleanedFilename,
	}
	store.Load()
	return &store
}

// func NewStoreFromSites(filename string, sites ...*Site) *Store {
// 	store := Store{
// 		FileName: filename,
// 		Sites:    sites,
// 	}
// 	return &store
// }

func (s *Store) AddToAvailableFiles(pathComponents ...string) {
	cleanPath, _ := tidyPath(pathComponents...)
	if DoesPathExist(cleanPath) {
		s.availableFiles = append(s.availableFiles, cleanPath)
	}
}

func (s *Store) verifyDefaultFiles() {
	for _, pathComponents := range DefaultStoreFileLocations {
		s.AddToAvailableFiles(pathComponents...)
	}
}

func (s *Store) Load() {
	rawContents, err := LoadFile(s.FileName)
	if nil != err {
		panic(err)
	}
	contents := strings.Split(string(rawContents), "\n")
	for _, url := range contents {
		if "" == url {
			continue
		}
		site := NewSite(url)
		if nil != site {
			s.Sites = append(s.Sites, site)
		}
	}
}
