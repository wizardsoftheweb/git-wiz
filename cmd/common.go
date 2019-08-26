package cmd

import (
	"github.com/spf13/cobra"
)

// Executes the command's help functionality
// This is useful for commands that have no base functionality and require
// an action or subcommand be run
func HelpOnly(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
