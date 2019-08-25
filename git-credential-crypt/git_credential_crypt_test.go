package main

import (
	"testing"

	"github.com/spf13/cobra"
	. "gopkg.in/check.v1"
)

func TestGitCredentialCrypt(t *testing.T) { TestingT(t) }

type GitCredentialCryptSuite struct {
	BaseSuite
}

var _ = Suite(&GitCredentialCryptSuite{})

func (s *GitCredentialCryptSuite) TestExecute(c *C) {
	var oldGitCredentialCryptCmd = &cobra.Command{}
	*oldGitCredentialCryptCmd = *GitCredentialCryptCmd
	dummy := func(cmd *cobra.Command, args []string) {}
	GitCredentialCryptCmd.SilenceErrors = true
	GitCredentialCryptCmd.DisableFlagParsing = true
	GitCredentialCryptCmd.PersistentPreRun = dummy
	GitCredentialCryptCmd.PreRun = dummy
	GitCredentialCryptCmd.Run = dummy
	GitCredentialCryptCmd.PostRun = dummy
	GitCredentialCryptCmd.PersistentPostRun = dummy
	err := Execute()
	c.Assert(err, IsNil)
	*GitCredentialCryptCmd = *oldGitCredentialCryptCmd
}
