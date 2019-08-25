package main

import (
	. "gopkg.in/check.v1"
)

type StoreSuite struct {
	BaseSuite
	store *Store
}

var _ = Suite(&StoreSuite{})

func (s *StoreSuite) SetUpTest(c *C) {
	s.store = &Store{}
}

func (s *StoreSuite) TearDownTest(c *C) {
}

func (s *StoreSuite) TestNewFromDiskDne(c *C) {
	s.store.FileName = s.currentFilename[1:]
	c.Assert(func() { s.store.Load() }, PanicMatches, "*no such file or directory")
}
