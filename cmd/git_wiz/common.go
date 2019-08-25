package git_wiz

import (
	"github.com/spf13/cobra"
)

func HelpOnly(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
