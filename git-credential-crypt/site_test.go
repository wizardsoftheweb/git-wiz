package main

import (
	"fmt"

	. "gopkg.in/check.v1"
)

type SiteSuite struct {
	BaseSuite
	defaultValue  string
	nothingActive [SiteNumberOfProperties]bool
	allActive     [SiteNumberOfProperties]bool
	emptyQuery    [SiteNumberOfProperties]string
	fullQuery     [SiteNumberOfProperties]string
	site          *Site
}

var _ = Suite(&SiteSuite{})

func (s *SiteSuite) SetUpTest(c *C) {
	s.site = &Site{}
}

func (s *SiteSuite) TearDownTest(c *C) {
}

func (s *SiteSuite) TestNewSiteValid(c *C) {
	input := []string{"http", "user", "password", "host", "path"}
	url := fmt.Sprintf(
		"%s://%s:%s@%s/%s",
		input[PositionSiteProtocol],
		input[PositionSiteUsername],
		input[PositionSitePassword],
		input[PositionSiteHost],
		input[PositionSitePath],
	)
	fmt.Println(url)
	site := NewSite(url)
	c.Assert(site.Protocol, Equals, input[PositionSiteProtocol])
	c.Assert(site.Username, Equals, input[PositionSiteUsername])
	c.Assert(site.Password, Equals, input[PositionSitePassword])
	c.Assert(site.Host, Equals, input[PositionSiteHost])
	c.Assert(site.Path, Equals, input[PositionSitePath])
}

func (s *SiteSuite) TestNewSiteInvalid(c *C) {
	site := NewSite("this won't work at all")
	c.Assert(site, IsNil)
}

func (s *SiteSuite) TestIsAMatchAllPermutations(c *C) {
	for _, entry := range siteSearchTestMatrix {
		s.site.sliceForSearch = entry.siteValues
		c.Assert(
			s.site.IsAMatch(entry.activated, entry.query),
			Equals,
			entry.result,
		)
	}
}

func (s *SiteSuite) TestToUrl(c *C) {
	matrix := []struct {
		urlComponents [SiteNumberOfProperties]string
		finalUrl      string
	}{
		{
			[SiteNumberOfProperties]string{"https", "rick", "james", "couch.com/"},
			"https://rick:james@couch.com",
		},
		{
			[SiteNumberOfProperties]string{"http", "user", "pass", "_L.;]0s:}", "!(<8B"},
			"http://user:pass@_L.%3B%5D0s%3A%7D%2F%21%28%3C8B",
		},
	}
	for _, entry := range matrix {
		s.site.Protocol = entry.urlComponents[PositionSiteProtocol]
		s.site.Username = entry.urlComponents[PositionSiteUsername]
		s.site.Password = entry.urlComponents[PositionSitePassword]
		s.site.Host = entry.urlComponents[PositionSiteHost]
		s.site.Path = entry.urlComponents[PositionSitePath]
		c.Assert(s.site.ToUrl(), Equals, entry.finalUrl)
	}

}

func (s *SiteSuite) TestDoesItWork(c *C) {
	site := Site{}
	c.Assert(site.isItUsable(), Equals, false)
	site.Protocol = "https"
	c.Assert(site.isItUsable(), Equals, false)
	site.Username = "rick"
	c.Assert(site.isItUsable(), Equals, false)
	site.Password = "james"
	c.Assert(site.isItUsable(), Equals, false)
	site.Host = "cool.beans"
	c.Assert(site.isItUsable(), Equals, true)
}

func (s *SiteSuite) TestIsPathOn(c *C) {
	s.site.Protocol = "tcp"
	c.Assert(s.site.isPathOn(), Equals, false)
}
