package cmd

import (
	"os"
	"runtime"
	"testing"

	"github.com/spf13/cobra"
	. "gopkg.in/check.v1"
)

func TestCommon(t *testing.T) { TestingT(t) }

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
	s.command = &cobra.Command{}
}

func (s *BaseSuite) TearDownSuite(c *C) {
	_ = os.Chdir(s.currentWorkingDirectory)
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
