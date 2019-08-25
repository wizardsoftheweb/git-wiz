package main

import (
	"regexp"
)

var PatternUrl = regexp.MustCompile(`^\s*(https?)(://)?([^/]+)/?(.+)?\s*$`)

type Site struct {
	Protocol string
	Host     string
	Path     string
	Username string
	Password string
	Url      string
}
