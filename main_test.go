package main

import (
	"testing"

	"github.com/spf13/cobra"
	. "gopkg.in/check.v1"

	"github.com/wizardsoftheweb/git-wiz/cmd"
)

func TestRootMain(t *testing.T) { TestingT(t) }

type MainSuite struct {
	BaseSuite
}

var _ = Suite(&MainSuite{})

func (s *MainSuite) TestMain(c *C) {
	var oldGitWizCmd = &cobra.Command{}
	*oldGitWizCmd = *cmd.GitWizCmd
	dummy := func(cmd *cobra.Command, args []string) {}
	cmd.GitWizCmd.SilenceErrors = true
	cmd.GitWizCmd.DisableFlagParsing = true
	cmd.GitWizCmd.PersistentPreRun = dummy
	cmd.GitWizCmd.PreRun = dummy
	cmd.GitWizCmd.Run = dummy
	cmd.GitWizCmd.PostRun = dummy
	cmd.GitWizCmd.PersistentPostRun = dummy
	c.Assert(
		func() {
			main()
		},
		Not(PanicMatches),
		"*",
	)
	*cmd.GitWizCmd = *oldGitWizCmd
}
