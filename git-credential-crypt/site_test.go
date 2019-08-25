package main

import (
	"fmt"

	. "gopkg.in/check.v1"
)

type SiteSuite struct {
	BaseSuite
}

var _ = Suite(&SiteSuite{})

func (s *SiteSuite) SetUpTest(c *C) {
}

func (s *SiteSuite) TearDownTest(c *C) {
}

func (s *SiteSuite) TestNewSite(c *C) {
	input := []string{"http", "user", "password", "host"}
	url := fmt.Sprintf(
		"%s://%s:%s@%s",
		input[PositionSiteProtocol],
		input[PositionSiteUsername],
		input[PositionSitePassword],
		input[PositionSiteHost],
	)
	site := NewSite(url)
	c.Assert(site.Protocol, Equals, input[PositionSiteProtocol])
	c.Assert(site.Username, Equals, input[PositionSiteUsername])
	c.Assert(site.Password, Equals, input[PositionSitePassword])
	c.Assert(site.Host, Equals, input[PositionSiteHost])
}
