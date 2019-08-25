package main

import (
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func TestSearchSite(t *testing.T) { TestingT(t) }

type SearchSiteSuite struct {
	BaseSuite
}

var _ = Suite(&SearchSiteSuite{})

func (s *SearchSiteSuite) SetUpTest(c *C) {
}

func (s *SearchSiteSuite) TearDownTest(c *C) {
}

func (s *SearchSiteSuite) TestSearchInput(c *C) {
	matrix := []struct {
		input      []string
		finalValue map[string]string
	}{
		{
			[]string{
				"protocol = https",
				"username = rick",
				"password = waste",
				"host = james",
				"path = couch",
			},
			map[string]string{
				"protocol": "https",
				"username": "rick",
				"host":     "james",
				"path":     "couch",
			},
		},
	}
	for _, entry := range matrix {
		input := strings.Join(entry.input, "\n")
		results := ParseSearchInput(input)
		c.Assert(results, DeepEquals, entry.finalValue)
	}
}
