package main

import (
	"os"
	"strings"
)

type Store struct {
	FileName     string
	Sites        []*Site
	defaultFiles []string
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

func (s *Store) defineDefaultFiles() {
	primary, _ := tidyPath("~", ".git-credentials")
	if DoesPathExist(primary) {
		s.defaultFiles = append(s.defaultFiles, primary)
	}
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	var secondary string
	if "" != xdgConfigHome {
		secondary, _ = tidyPath(xdgConfigHome, "git", "credentials")
		s.defaultFiles = append(s.defaultFiles, secondary)
	}
	if DoesPathExist(secondary) {
		s.defaultFiles = append(s.defaultFiles, secondary)
	}
	tertiary, _ := tidyPath("~", ".config", "git", "credentials")
	s.defaultFiles = append(s.defaultFiles, tertiary)
	if DoesPathExist(tertiary) {
		s.defaultFiles = append(s.defaultFiles, tertiary)
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
		s.Sites = append(s.Sites, NewSite(url))
	}
}
