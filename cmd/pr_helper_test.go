package cmd

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestPrHelper(t *testing.T) { TestingT(t) }

type PrHelperSuite struct {
	BaseSuite
}

var _ = Suite(&PrHelperSuite{})

func (s *PrHelperSuite) SetUpTest(c *C) {
}

func (s *PrHelperSuite) TearDownTest(c *C) {
}
