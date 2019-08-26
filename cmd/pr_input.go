package cmd

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/chzyer/readline"
)

// Passing in a string gets a cow saying that string
func prepCow(words string) string {
	say, _ := cowsay.Say(
		cowsay.Phrase(words),
		cowsay.Type("default"),
	)
	return say
}

// Use the discovered local information to suggest the contents of a PR payload.
func compileSuggestedPrBody() *GitHubPrRequest {
	discovery := CompletePrDiscovery()
	return &GitHubPrRequest{
		Title: discovery.assumedCurrentBranch,
		Base:  discovery.suggestedBase,
		Head:  discovery.assumedCurrentBranch,
		Body: "The tool that was used to automatically generate this PR " +
			"doesn't do bodies yet and the author of this has neglected to " +
			"update this via the GUI.",
		MaintainerCanModify: true,
	}
}

// This function is used by all user approvals that only require one line of
// input. It prints instructions to the user, prefills the entry with the
// provided suuggestion, and returns the approved string.
func approveOneLineItem(instructions, promptTitle, suggestedItem string) string {
	fmt.Println(prepCow(instructions))
	input, _ := readline.New(fmt.Sprintf("%s> ", promptTitle))
	_, _ = input.WriteStdin([]byte(suggestedItem))
	result, _ := input.Readline()
	return result
}

// Get user approval for the PR title
func approvePrTitle(suggestedTitle string) string {
	return approveOneLineItem(
		"The name of your current branch was chosen "+
			"as a title for this PR. If you'd like the PR to have a different "+
			"title, please enter one now and hit return when you're finished.",
		"Title",
		suggestedTitle,
	)
}

// Get user approval for the PR base branch (where it's going)
func approvePrBase(suggestedBase string) string {
	return approveOneLineItem(
		"If you're using GitFlow, your prefix branch's base "+
			"was selected as the PR base. If you're not using GitFlow you'll "+
			"have to do manual discovery like a barbarian",
		"Base Branch",
		suggestedBase,
	)
}

// Get user approval for the PR head branch (where it's coming from)
func approvePrHead(suggestedHead string) string {
	return approveOneLineItem(
		"The branch you're currently working on was selected as the "+
			"base branch for the PR. If you need to change that, please "+
			"enter a new name and hit return when you're finished.",
		"Head Branch",
		suggestedHead,
	)
}

// Create a description of the PR
func createPrBody(suggestedBody string) string {
	return approveOneLineItem(
		"Please enter a short description of this PR.\n\nUntil wiz "+
			"is built to load from files, the GH GUI is still, "+
			"unfortunately, the best place to draft your PR body.\n\n"+
			"Hitting return ends the body so keep it short. ¯\\_(ツ)_/¯",
		"Body",
		suggestedBody,
	)
}

// Present all suggestions to the user one at a time. It checks at the end to
// see if any changes need to be made. If there are, the function loops back
// over itself.
func loopUntilPrItemsAreApproved(request *GitHubPrRequest) *GitHubPrRequest {
	request.Title = approvePrTitle(request.Title)
	request.Base = approvePrBase(request.Base)
	request.Head = approvePrHead(request.Head)
	request.Body = createPrBody(request.Body)
	approval := approveOneLineItem(
		"If you need to make any changes, enter 'yes'. Anything "+
			"other than 'yes' is interpreted as approval and will POST this "+
			"PR request to the repo.",
		"'yes' to make changes",
		"",
	)
	if "yes" == approval {
		return loopUntilPrItemsAreApproved(request)
	}
	return request
}

// This is unused
// It will compile the payload for requesting reviewers on the PR
func compileSuggestedReviewBody(owner, repo string) *GithubReviewRequest {
	reviewers := CompleteReviewDiscovery(owner, repo)
	return &GithubReviewRequest{
		Reviewers: reviewers,
	}
}
