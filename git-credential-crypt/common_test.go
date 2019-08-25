package main

import (
	"errors"
	"os"
	"runtime"
	"testing"

	"github.com/spf13/cobra"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BaseSuite struct {
	workingDirectory        string
	currentFilename         string
	currentWorkingDirectory string
	errorMessage            string
	command                 *cobra.Command
	args                    []string
}

var _ = Suite(&BaseSuite{})

func (s *BaseSuite) SetUpSuite(c *C) {
	s.currentWorkingDirectory, _ = os.Getwd()
	s.workingDirectory = c.MkDir()
	_ = os.Chdir(s.workingDirectory)
	_, s.currentFilename, _, _ = runtime.Caller(0)
	s.errorMessage = "shared file error"
	s.command = &cobra.Command{}
}

func (s *BaseSuite) TearDownSuite(c *C) {
	_ = os.Chdir(s.currentWorkingDirectory)
}
func (s *BaseSuite) brokenPathTidier(input ...string) (string, error) {
	return "", errors.New(s.errorMessage)
}

type CommonSuite struct {
	BaseSuite
}

var _ = Suite(&CommonSuite{})

func (s *CommonSuite) SetUpTest(c *C) {
}

func (s *CommonSuite) TearDownTest(c *C) {
}

func (s *CommonSuite) TestHelpOnly(c *C) {
	c.Assert(
		func() {
			HelpOnly(s.command, s.args)
		},
		Not(PanicMatches),
		"*",
	)
}
