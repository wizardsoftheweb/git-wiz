package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var PatternSiteEntry = regexp.MustCompile(`^\s*(.*)://(.*):(.*)@([^/\n]*)/?(.*)\s*$`)

type MatchPositionSite int

const SiteNumberOfProperties = 5

const (
	PositionSiteProtocol MatchPositionSite = iota
	PositionSiteUsername
	PositionSitePassword
	PositionSiteHost
	PositionSitePath
)

type Site struct {
	Protocol       string
	Username       string
	Password       string
	Host           string
	Path           string
	Url            string
	sliceForSearch [SiteNumberOfProperties]string
}

func NewSite(url string) *Site {
	components := ExplodeUrl(url)
	if 0 < len(components) {
		site := Site{Url: url}
		site.parseUrl(components)
		return &site
	}
	return nil
}

func (s *Site) isPathOn() bool {
	return strings.Contains(s.Protocol, "http") && DoPathsMatter()
}

func (s *Site) decodeComponent(value string) string {
	decoded, _ := url.QueryUnescape(value)
	return decoded
}

func ExplodeUrl(workingUrl string) []string {
	matches := PatternSiteEntry.FindAllStringSubmatch(workingUrl, -1)
	if 1 > len(matches) {
		return []string{}
	}
	return matches[0]
}

func (s *Site) doesItWork() bool {
	return "" != s.Protocol &&
		("" != s.Host || "" != s.Path) &&
		"" != s.Username &&
		"" != s.Password
}

func (s *Site) updateSliceForSearch() {
	s.sliceForSearch = [SiteNumberOfProperties]string{
		s.Protocol,
		s.Username,
		s.Password,
		s.Host,
		s.Path,
	}
}

func (s *Site) parseUrl(components []string) {
	s.Protocol = s.decodeComponent(components[PositionSiteProtocol+1])
	s.Username = s.decodeComponent(components[PositionSiteUsername+1])
	s.Password = s.decodeComponent(components[PositionSitePassword+1])
	s.Host = strings.TrimSuffix(
		s.decodeComponent(components[PositionSiteHost+1]),
		"/",
	)
	s.Path = strings.TrimSuffix(
		s.decodeComponent(components[PositionSitePath+1]),
		"/",
	)
	s.updateSliceForSearch()
}

func (s *Site) IsAMatch(activated [SiteNumberOfProperties]bool, query [SiteNumberOfProperties]string) bool {
	for index, active := range activated {
		if int(PositionSitePassword) == index || (int(PositionSitePath) == index && !DoPathsMatter()) {
			continue
		}
		if active {
			if "" == query[index] || query[index] != s.sliceForSearch[index] {
				return false
			}
		}
	}
	return true
}

func (s *Site) ToUrl() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s/%s",
		s.Protocol,
		url.QueryEscape(s.Username),
		url.QueryEscape(s.Password),
		url.QueryEscape(s.Host),
		url.QueryEscape(s.Path),
	)
}
