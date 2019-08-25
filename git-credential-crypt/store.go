package main

import (
	"os"
	"strings"
)

var UrlComponents = []string{"protocol", "username", "password", "host", "path"}

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
	cleanedFilename, _ := pathTidier(filename)
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
	cleanPath, _ := pathTidier(pathComponents...)
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

func (s *Store) Write() {
	s.FileName, _ = pathTidier(s.FileName)
	var siteLines []string
	for _, site := range s.Sites {
		if site.isItUsable() {
			siteLines = append(siteLines, site.ToUrl())
		}
	}
	blob := strings.Join(siteLines, "\n")
	err := WriteFile([]byte(blob+"\n"), 0600, s.FileName)
	if nil != err {
		panic(err)
	}
}

func (s *Store) constructSearchParameters(incoming map[string]string) ([SiteNumberOfProperties]bool, [SiteNumberOfProperties]string) {
	activated := [SiteNumberOfProperties]bool{}
	query := [SiteNumberOfProperties]string{}
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
	query[PositionSitePath] = strings.TrimSuffix(query[PositionSitePath], "/")
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
