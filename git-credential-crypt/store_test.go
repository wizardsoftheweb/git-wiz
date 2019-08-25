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

func (s *StoreSuite) TestNewFromDisk(c *C) {
	store := NewStoreFromDisk(s.credsPath)
	c.Assert(0 < len(store.Sites), Equals, true)

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

func (s *StoreSuite) TestAddToAvailableFilesNotToday(c *C) {
	c.Assert(s.store.availableFiles, HasLen, 0)
	s.store.AddToAvailableFiles(s.currentFilename[1:])
	c.Assert(s.store.availableFiles, HasLen, 0)
}

func (s *StoreSuite) TestAddToAvailableFilesExists(c *C) {
	c.Assert(s.store.availableFiles, HasLen, 0)
	s.store.AddToAvailableFiles(s.currentFilename)
	c.Assert(s.store.availableFiles, HasLen, 1)
}

func (s *StoreSuite) TestVerifyDefaultFiles(c *C) {
	c.Assert(
		s.store.verifyDefaultFiles,
		Not(Panics),
		"*",
	)
}
