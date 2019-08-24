package main

import (
	"os"
	"runtime"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BaseSuite struct {
	workingDirectory string
	currentFilename string
	currentWorkingDirectory string
}

var _ = Suite(&BaseSuite{})

func (s *BaseSuite) SetUpSuite(c *C) {
	s.currentWorkingDirectory, _ = os.Getwd()
	s.workingDirectory = c.MkDir()
	_ = os.Chdir(s.workingDirectory)
	_, s.currentFilename, _, _ = runtime.Caller(0)
}

func (s *BaseSuite) TearDownSuite(c *C) {
	_ = os.Chdir(s.currentWorkingDirectory)
}
