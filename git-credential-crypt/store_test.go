package main

import (
	"path/filepath"

	. "gopkg.in/check.v1"
)

type StoreSuite struct {
	BaseSuite
	store     *Store
	credsPath string
}

var _ = Suite(&StoreSuite{})

func (s *StoreSuite) SetUpTest(c *C) {
	s.store = &Store{
		Sites: []*Site{},
	}
	s.credsPath = filepath.Join(filepath.Dir(s.currentFilename), "docs", "sample-git-credentials")
}

func (s *StoreSuite) TearDownTest(c *C) {
}

func (s *StoreSuite) TestLoadDne(c *C) {
	s.store.FileName = s.currentFilename[1:]
	c.Assert(func() { s.store.Load() }, PanicMatches, "*no such file or directory")
}

func (s *StoreSuite) TestLoadExists(c *C) {
	s.store.FileName = s.credsPath
	c.Assert(s.store.Sites, HasLen, 0)
	s.store.Load()
	c.Assert(0 < len(s.store.Sites), Equals, true)

}
