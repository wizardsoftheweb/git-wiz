package git_wiz

import (
	"github.com/spf13/cobra"
	. "gopkg.in/check.v1"
)

type GitWizSuite struct {
	BaseSuite
}

var _ = Suite(&GitWizSuite{})

func (s *GitWizSuite) TestExecute(c *C) {
	var oldGitWizCmd = &cobra.Command{}
	*oldGitWizCmd = *GitWizCmd
	dummy := func(cmd *cobra.Command, args []string) {}
	GitWizCmd.SilenceErrors = true
	GitWizCmd.DisableFlagParsing = true
	GitWizCmd.PersistentPreRun = dummy
	GitWizCmd.PreRun = dummy
	GitWizCmd.Run = dummy
	GitWizCmd.PostRun = dummy
	GitWizCmd.PersistentPostRun = dummy
	err := Execute()
	c.Assert(err, IsNil)
	*GitWizCmd = *oldGitWizCmd
}
