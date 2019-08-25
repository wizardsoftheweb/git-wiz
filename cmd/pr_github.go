package cmd

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
type ReviewRequest struct {
	// An array of user logins that will be requested.
	Reviewers []string `json:"reviewers"`
	// An array of team slugs that will be requested.
	TeamReviewers []string `json:"reviewers"`
}

// https://developer.github.com/v3/repos/#parameters-6
// These docs are courtesy of the API documentation.
type ListContributorsRequest struct {
	// Set to 1 or true to include anonymous contributors in results.
	Anon string `json:"anon"`
}
