package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestGitConfig(t *testing.T) { TestingT(t) }

type GitConfigSuite struct {
	BaseSuite
}

var _ = Suite(&GitConfigSuite{})

func (s *GitConfigSuite) SetUpTest(c *C) {
}

func (s *GitConfigSuite) TearDownTest(c *C) {
}
