package main

import (
	"net/url"
	"regexp"
)

var PatternSiteEntry = regexp.MustCompile(`^\s*(.*)://(.*):(.*)@(.*)\s*$`)

type MatchPositionSite int

const (
	PositionSiteProtocol MatchPositionSite = iota + 1
	PositionSiteUsername
	PositionSitePassword
	PositionSiteHost
)

type Site struct {
	Protocol string
	Host     string
	Username string
	Password string
	Url      string
}

func NewSite(url string) *Site {
	site := Site{Url: url}
	site.parseUrl()
	return &site
}

func (s *Site) decodeComponent(value string) string {
	decoded, _ := url.QueryUnescape(value)
	return decoded
}

func (s *Site) parseUrl() {
	matches := PatternSiteEntry.FindAllStringSubmatch(s.Url, -1)[0]
	s.Protocol = s.decodeComponent(matches[PositionSiteProtocol])
	s.Username = s.decodeComponent(matches[PositionSiteUsername])
	s.Password = s.decodeComponent(matches[PositionSitePassword])
	s.Host = s.decodeComponent(matches[PositionSiteHost])
}
