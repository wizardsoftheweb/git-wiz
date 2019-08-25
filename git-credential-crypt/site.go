package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var PatternSiteEntry = regexp.MustCompile(`^\s*(.*)://(.*)\s*$`)

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
	fmt.Println(components)
	if 0 < len(components) {
		site := Site{Url: url}
		site.parseUrl(components)
		return &site
	}
	return nil
}

func (s *Site) isPathOn() bool {
	return DoPathsMatter() && ("https" == s.Protocol || "http" == s.Protocol)
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
	return matches[0][1:]
}

func (s *Site) isItUsable() bool {
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
	s.Protocol = s.decodeComponent(components[0])
	explodedTail := strings.SplitN(components[1], "@", 2)
	if 2 == len(explodedTail) {
		frontExplosion := strings.SplitN(explodedTail[0], ":", 2)
		s.Username = s.decodeComponent(frontExplosion[0])
		if 2 == len(frontExplosion) {
			s.Password = s.decodeComponent(frontExplosion[1])
		}
		backExplosion := strings.SplitN(explodedTail[1], "/", 2)
		s.Host = s.decodeComponent(backExplosion[0])
		if 2 == len(backExplosion) {
			s.Path = s.decodeComponent(backExplosion[1])
		}
	} else {
		s.Host = explodedTail[0]
	}
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
	pathComponent := ""
	if DoPathsMatter() && "" != s.Path {
		pathComponent = "/" + s.Path
	}
	return strings.TrimSuffix(
		fmt.Sprintf(
			"%s://%s:%s@%s%s",
			s.Protocol,
			url.QueryEscape(s.Username),
			url.QueryEscape(s.Password),
			url.QueryEscape(s.Host),
			url.QueryEscape(pathComponent),
		),
		"%2F",
	)
}
