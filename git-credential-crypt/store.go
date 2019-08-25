package main

import (
	"os"
	"strings"
)

var UrlComponents = []string{"protocol", "username", "password", "host"}

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

func (s *Store) constructSearchParameters(incoming map[string]string) ([4]bool, [4]string) {
	activated := [4]bool{}
	query := [4]string{}
	for index, value := range UrlComponents {
		input, ok := incoming[value]
		activated[index] = ok
		if ok {
			query[index] = input
		} else {
			query[index] = ""
		}
	}
	query[PositionSiteHost] = strings.TrimSuffix(query[PositionSiteHost], "/")
	return activated, query
}

func (s *Store) Get(incoming map[string]string) *Site {
	activated, query := s.constructSearchParameters(incoming)
	for _, site := range s.Sites {
		if site.IsAMatch(activated, query) {
			return site
		}
	}
	return nil
}
