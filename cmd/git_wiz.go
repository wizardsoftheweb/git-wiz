package cmd

import (
	"github.com/spf13/cobra"
)

// The is the git-wiz version only. Other components may have
// different versions.
var PackageVersion = "undefined"

// This is the flag to check the command version
var ShowVersion bool

// The verbosity flag is a count flag, ie the more there are the more verbose
// it gets.
var VerbosityFlagValue int

func init() {
	GitWizCmd.PersistentFlags().BoolVarP(
		&ShowVersion,
		"version",
		"V",
		false,
		"Prints the app version",
	)
	GitWizCmd.PersistentFlags().CountVarP(
		&VerbosityFlagValue,
		"verbose",
		"v",
		"Increases application verbosity",
	)
}

// This is the primary cmd runner and exposes git-wiz
func Execute() error {
	return GitWizCmd.Execute()
}

// git-wiz has no base functionality. It must be used with subcommands.
var GitWizCmd = &cobra.Command{
	Use:   "git-wiz",
	Short: "i have no idea what im doing",
	Long: "wiz addresses some QoL issues I have with some of the tools I " +
		"use, provides an easier medium to consume some of my git " +
		"experiments, and might hopefully provide value to someone else. I " +
		"have my doubts about that last claim.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ShowVersion {
			cmd.Printf("%x version %s\n", cmd.Use, PackageVersion)
			cmd.TraverseChildren = false
			cmd.Run = func(cmd *cobra.Command, args []string) {}
		}
	},
	Run: HelpOnly,
}

// This is a catch-all error handler that kills the program when an
// error occurs.
func whereErrorsGoToDie(err error) {
	if nil != err {
		panic(err)
	}
}
