package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// prCmd generates a PR using information about the current directory state.
// All of its content goes through final approval with the user before being
// posted to GH.
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "PRs through the CLI where your git flow already is",
	Long:  "Just GH for now. This may or may not ever be finished.",
	Run: func(cmd *cobra.Command, args []string) {
		rawPr := compileSuggestedPrBody()
		approvedPr := loopUntilPrItemsAreApproved(rawPr)
		repoOwner := os.Getenv(EnvVariableThatHoldsMyRepoOwner)
		repoName := os.Getenv(EnvVariableThatHoldsMyRepoName)
		prRequestBody, _ := json.Marshal(approvedPr)
		prResponse := createPullRequest(repoOwner, repoName, prRequestBody)
		fmt.Println(string(prResponse))
	},
}

func init() {
	GitWizCmd.AddCommand(prCmd)
}
