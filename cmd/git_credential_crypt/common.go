package git_credential_crypt

import (
	"github.com/spf13/cobra"
)

func HelpOnly(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
