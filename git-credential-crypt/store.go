package main

import (
	"strings"
)

type Store struct {
	FileName string
	Sites    []*Site
}

func NewStoreFromDisk(filename string) *Store {
	store := Store{
		FileName: filename,
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
