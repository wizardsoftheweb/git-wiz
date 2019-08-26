package cmd

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/chzyer/readline"
)

func prepCow(words string) string {
	say, _ := cowsay.Say(
		cowsay.Phrase(words),
		cowsay.Type("default"),
	)
	return say
}

func compileSuggestedPrBody() *GitHubPrRequest {
	discovery := CompletePrDiscovery()
	return &GitHubPrRequest{
		Title: discovery.assumedCurrentBranch,
		Base:  discovery.assumedCurrentBranch,
		Head:  discovery.suggestedHead,
		Body: "The tool that was used to automatically generate this PR " +
			"doesn't do bodies yet and the author of this has neglected to " +
			"update this via the GUI.",
		MaintainerCanModify: true,
	}
}

func approveOneLineItem(instructions, promptTitle, suggestedItem string) string {
	fmt.Println(prepCow(instructions))
	input, _ := readline.New(fmt.Sprintf("%s> ", promptTitle))
	_, _ = input.WriteStdin([]byte(suggestedItem))
	result, _ := input.Readline()
	return result
}

func approvePrTitle(suggestedTitle string) string {
	return approveOneLineItem(
		"The name of your current branch was chosen "+
			"as a title for this PR. If you'd like the PR to have a different "+
			"title, please enter one now and hit return when you're finished.",
		"Title",
		suggestedTitle,
	)
}

func approvePrBase(suggestedBase string) string {
	return approveOneLineItem(
		"The branch you're currently working on was selected as the "+
			"base branch for the PR. If you need to change that, please "+
			"enter a new name and hit return when you're finished.",
		"Base Branch",
		suggestedBase,
	)
}

func approvePrHead(suggestedHead string) string {
	return approveOneLineItem(
		"If you're using GitFlow, your prefix branch's base "+
			"was selected as the PR head. If you're not using GitFlow you'll "+
			"have to do manual discovery like a barbarian",
		"Head Branch",
		suggestedHead,
	)
}

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

func sharePr(pr *GitHubPrRequest) {
	fmt.Println(pr)
}

func compileSuggestedReviewBody(owner, repo string) *GithubReviewRequest {
	reviewers := CompleteReviewDiscovery(owner, repo)
	return &GithubReviewRequest{
		Reviewers: reviewers,
	}
}
