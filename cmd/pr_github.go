package cmd

import (
	"fmt"
)

// https://developer.github.com/v3/pulls/#input
// These docs are courtesy of the API documentation.
type GitHubPrRequest struct {
	// Required. The title of the pull request.
	Title string `json:"title"`
	// Required. The name of the branch where your changes are implemented. For
	// cross-repository pull requests in the same network, namespace head with
	// a user like this: username:branch.
	Head string `json:"head"`
	// Required. The name of the branch you want the changes pulled into. This
	// should be an existing branch on the current repository. You cannot submit
	// a pull request to one repository that requests a merge to a base of
	// another repository.
	Base string `json:"base"`
	// The contents of the pull request.
	Body string `json:"body"`
	// Indicates whether maintainers can modify the pull request.
	MaintainerCanModify bool `json:"maintainer_can_modify"`
	// Indicates whether the pull request is a draft.
	// I'm not using this yet because it's a preview feature and I don't want
	// to mess with that.
	Draft bool `json:"draft"`
}

// https://developer.github.com/v3/pulls/review_requests/#input
// These docs are courtesy of the API documentation.
type GithubReviewRequest struct {
	// An array of user logins that will be requested.
	Reviewers []string `json:"reviewers"`
	// An array of team slugs that will be requested.
	TeamReviewers []string `json:"reviewers"`
}

// This sends a request to the following endpoint to discover users that have
// collaborated on this repo
//
// 		GET  /repos/:owner/:repo/collaborators
func getCollaboratorList(owner, repo string) []byte {
	resource := fmt.Sprintf("repos/%s/%s/collaborators", owner, repo)
	return getResource(resource, nil)
}

// This sends a request to the following endpoint to create a PR
//
// 		POST /repos/:owner/:repo/pulls
func createPullRequest(owner, repo string, requestBody []byte) []byte {
	resource := fmt.Sprintf("repos/%s/%s/pulls", owner, repo)
	return getResource(resource, requestBody)
}

// This sends a request to the following endpoint to request a list of reviewers
// on the PR
//
// 		// POST /repos/:owner/:repo/pulls/:pull_number/reviews
func requestPrReview(owner, repo, pullNumber string, requestBody []byte) []byte {
	resource := fmt.Sprintf("repos/%s/%s/pulls/%s/reviews", owner, repo, pullNumber)
	return getResource(resource, requestBody)
}
