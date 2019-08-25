package main

import (
	"fmt"

	. "gopkg.in/check.v1"
)

type SiteSuite struct {
	BaseSuite
	defaultValue  string
	nothingActive [4]bool
	allActive     [4]bool
	emptyQuery    [4]string
	fullQuery     [4]string
	site          *Site
}

var _ = Suite(&SiteSuite{})

func (s *SiteSuite) SetUpTest(c *C) {
	s.defaultValue = "qqq"
	s.nothingActive = [4]bool{false, false, false, false}
	s.allActive = [4]bool{true, true, true, true}
	s.emptyQuery = [4]string{"", "", "", ""}
	s.fullQuery = [4]string{s.defaultValue, s.defaultValue, s.defaultValue, s.defaultValue}
	s.site = &Site{}
}

func (s *SiteSuite) TearDownTest(c *C) {
}

func (s *SiteSuite) TestNewSiteValid(c *C) {
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
		urlComponents [4]string
		finalUrl      string
	}{
		{
			[4]string{"https", "rick", "james", "couch.com/"},
			"https://rick:james@couch.com%2F",
		},
		{
			[4]string{"http", "user", "pass", "_L.;]0s:}/!(<8B"},
			"http://user:pass@_L.%3B%5D0s%3A%7D%2F%21%28%3C8B",
		},
	}
	for _, entry := range matrix {
		s.site.Protocol = entry.urlComponents[PositionSiteProtocol]
		s.site.Username = entry.urlComponents[PositionSiteUsername]
		s.site.Password = entry.urlComponents[PositionSitePassword]
		s.site.Host = entry.urlComponents[PositionSiteHost]
		c.Assert(s.site.ToUrl(), Equals, entry.finalUrl)
	}

}
