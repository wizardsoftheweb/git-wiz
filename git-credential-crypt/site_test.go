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
		input[PositionSiteProtocol-1],
		input[PositionSiteUsername-1],
		input[PositionSitePassword-1],
		input[PositionSiteHost-1],
	)
	site := NewSite(url)
	c.Assert(site.Protocol, Equals, input[PositionSiteProtocol-1])
	c.Assert(site.Username, Equals, input[PositionSiteUsername-1])
	c.Assert(site.Password, Equals, input[PositionSitePassword-1])
	c.Assert(site.Host, Equals, input[PositionSiteHost-1])
}
