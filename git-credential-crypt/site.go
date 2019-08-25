package main

import (
	"net/url"
	"regexp"
	"strings"
)

var PatternSiteEntry = regexp.MustCompile(`^\s*(.*)://(.*):(.*)@(.*)/?\s*$`)

type MatchPositionSite int

const (
	PositionSiteProtocol MatchPositionSite = iota
	PositionSiteUsername
	PositionSitePassword
	PositionSiteHost
)

type Site struct {
	Protocol       string
	Host           string
	Username       string
	Password       string
	Url            string
	sliceForSearch [4]string
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

func (s *Site) parseUrl(components []string) {
	s.Protocol = s.decodeComponent(components[PositionSiteProtocol+1])
	s.Username = s.decodeComponent(components[PositionSiteUsername+1])
	s.Password = s.decodeComponent(components[PositionSitePassword+1])
	s.Host = strings.TrimSuffix(
		s.decodeComponent(components[PositionSiteHost+1]),
		"/",
	)
	s.sliceForSearch = [4]string{
		s.Protocol,
		s.Username,
		s.Password,
		s.Host,
	}
}

func (s *Site) IsAMatch(activated [4]bool, query [4]string) bool {
	for index, active := range activated {
		if int(PositionSitePassword) == index {
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
