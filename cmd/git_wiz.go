package cmd

import (
	"github.com/spf13/cobra"
)

var PackageVersion = "0.0.0"
var VerbosityFlagValue int

func init() {
	GitWizCmd.PersistentFlags().CountVarP(
		&VerbosityFlagValue,
		"verbose",
		"v",
		"Increases application verbosity",
	)
}

func Execute() error {
	return GitWizCmd.Execute()
}

var GitWizCmd = &cobra.Command{
	Use:     "git-wiz",
	Version: PackageVersion,
	Short:   "i have no idea what im doing",
	Run:     HelpOnly,
}
