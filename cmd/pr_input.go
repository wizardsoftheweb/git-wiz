package cmd

import (
	"fmt"
)

func compileSuggestedPrBody() *GitHubPrRequest {
	discovery := CompletePrDiscovery()
	return &GitHubPrRequest{
		Title:               discovery.assumedCurrentBranch,
		Base:                discovery.assumedCurrentBranch,
		Head:                discovery.suggestedHead,
		Body:                "",
		MaintainerCanModify: true,
	}
}

func approvePrTitle(suggestedTitle string) {

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
