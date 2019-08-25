package main

import (
	"strings"

	"github.com/prometheus/common/log"
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
		log.Fatal(err)
	}
	contents := strings.Split(string(rawContents), "\n")
	for index, url := range contents {
		s.Sites[index] = NewSite(url)
	}
}
