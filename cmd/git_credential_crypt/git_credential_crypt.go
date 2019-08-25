package git_credential_crypt

import (
	"github.com/spf13/cobra"
)

var PackageVersion = "0.0.0"
var VerbosityFlagValue int

func init() {
	GitCredentialCryptCmd.PersistentFlags().CountVarP(
		&VerbosityFlagValue,
		"verbose",
		"v",
		"Increases application verbosity",
	)
}

func Execute() error {
	return GitCredentialCryptCmd.Execute()
}

var GitCredentialCryptCmd = &cobra.Command{
	Use:     "git-credential-crypt",
	Version: PackageVersion,
	Short:   "An alternative, loosely encrypted solution to store and cache",
	Run:     HelpOnly,
}
