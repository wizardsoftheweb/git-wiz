package git_credentials_store

import (
	"github.com/spf13/cobra"
)

func HelpOnly(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
