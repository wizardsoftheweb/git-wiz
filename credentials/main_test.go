package main

import (
	. "gopkg.in/check.v1"
)

type MainSuite struct {
	BaseSuite
}

var _ = Suite(&MainSuite{})

func (s *MainSuite) SetUpTest(c *C) {
}

func (s *MainSuite) TearDownTest(c *C) {
}

func (s *MainSuite) TestMain(c *C) {
	c.Assert(
		func() {
			main()
		},
		Not(PanicMatches),
		"*",
	)
}
