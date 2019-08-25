package cmd

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestPrGithub(t *testing.T) { TestingT(t) }

type PrGithubSuite struct {
	BaseSuite
}

var _ = Suite(&PrGithubSuite{})

func (s *PrGithubSuite) SetUpTest(c *C) {
}

func (s *PrGithubSuite) TearDownTest(c *C) {
}
