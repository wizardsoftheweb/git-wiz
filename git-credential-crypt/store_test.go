package main

import (
	"path/filepath"
	"testing"

	. "gopkg.in/check.v1"
)

func TestStore(t *testing.T) { TestingT(t) }

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
	pathTidier = tidyPath
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

func (s *StoreSuite) TestConstructSearchParameters(c *C) {
	matrix := []struct {
		key        string
		value      string
		finalValue string
	}{
		{UrlComponents[PositionSiteProtocol], "https", "https"},
		{UrlComponents[PositionSiteUsername], "rick", "rick"},
		{UrlComponents[PositionSitePassword], "james", "james"},
		{UrlComponents[PositionSiteHost], "test.com/", "test.com"},
	}
	for index, entry := range matrix {
		incoming := map[string]string{
			entry.key: entry.value,
		}
		activated, query := s.store.constructSearchParameters(incoming)
		for counter := 0; counter < 4; counter++ {
			if index != counter {
				c.Assert(activated[counter], Equals, false)
				c.Assert(query[counter], Equals, "")
			} else {
				c.Assert(activated[counter], Equals, true)
				c.Assert(query[counter], Equals, entry.finalValue)

			}
		}
	}
}

func (s *StoreSuite) TestGetNothing(c *C) {
	incoming := make(map[string]string)
	c.Assert(s.store.Get(incoming), IsNil)
}

func (s *StoreSuite) TestGetSuccessful(c *C) {
	incoming := map[string]string{
		"protocol": "https",
		"host":     "host",
	}
	s.store.FileName = s.credsPath
	s.store.Load()
	site := s.store.Get(incoming)
	c.Assert(site.Protocol, Equals, "https")
	c.Assert(site.Username, Equals, "user")
	c.Assert(site.Password, Equals, "pass")
	c.Assert(site.Host, Equals, "host")
}

func (s *StoreSuite) TestWriteCorrectly(c *C) {
	s.store = NewStoreFromDisk(s.credsPath)
	s.store.FileName = filepath.Join(s.workingDirectory, "credentials")
	s.store.Write()
}

func (s *StoreSuite) TestWriteFailure(c *C) {
	s.store = NewStoreFromDisk(s.credsPath)
	pathTidier = s.brokenPathTidier
	s.store.FileName = s.currentFilename[1:]
	c.Assert(
		func() {
			s.store.Write()
		},
		PanicMatches,
		s.errorMessage,
	)
}
