package cmd

import (
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "PRs through the CLI where your git flow already is",
	Long:  "Just GH for now. This may or may not ever be finished.",
	Run: func(cmd *cobra.Command, args []string) {
		Demo = WotwPrRequest{}
		Demo.discoverGitConfigDirectory()
	},
}

func init() {
	GitWizCmd.AddCommand(prCmd)
}
