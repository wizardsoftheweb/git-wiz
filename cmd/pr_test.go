package cmd

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestPr(t *testing.T) { TestingT(t) }

type PrSuite struct {
	BaseSuite
}

var _ = Suite(&PrSuite{})

func (s *PrSuite) SetUpTest(c *C) {
}

func (s *PrSuite) TearDownTest(c *C) {
}
