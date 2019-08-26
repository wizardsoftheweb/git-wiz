package cmd

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// This is the name of an environment variable that contains a personal access
// that can be used to edit repos
const EnvVariableThatHoldsMyPat = "GH_DEV_PAT"

// This is the name of an environment variable that contains the username
// associated with the PAT
const EnvVariableThatHoldsMyGhUser = "GH_USERNAME"

// This is the name of an environment variable that contains the name of the
// the owner of the repo that will be PR'd into. This will eventually be input.
const EnvVariableThatHoldsMyRepoOwner = "GH_OWNER"

// This is the name of an environment variable that contains the name of the
// the repo that will be PR'd into. This will eventually be input.
const EnvVariableThatHoldsMyRepoName = "GH_REPO"

// This holds all the local info that will be used to suggest a payload for the
// API call to build the PR
type PrDiscovery struct {
	// If it's there, this becomes much easier.
	hasGitFlow bool
	// The root path of the repo that contains what I need.
	gitConfigDirectory string
	// This comes from the CLI before the app reaches out to the user so it
	// could be suspect.
	assumedCurrentBranch string
	// Theoretically tracked in git-flow
	suggestedBase string
}

// This holds all the local info that will be used to suggest a payload for the
// API call to ask for reviews
type ReviewDiscovery struct {
	// Contributors from local VCS
	localSuggestedReviewers []string
	// Contributors from remote VCS
	remoteSuggestedReviewers []string
}

var (
	// Hopefully one of these will actually work
	commandsToGetCurrentBranch = [][]string{
		{"git", "branch", "--show-current"},
		{"git", "rev-parse", "--abbrev-ref", "HEAD"},
		{"git", "symbolic-ref", "--short", "HEAD"},
	}
)

// If we're in the repo, rev-parse should get us there. If not, we'll have to
// prompt for it.
func (w *PrDiscovery) discoverGitConfigDirectory() {
	action := exec.Command("git", []string{"rev-parse", "--show-toplevel"}...)
	result, err := action.CombinedOutput()
	if nil != err {
		return
	}
	w.gitConfigDirectory = filepath.Join(string(result), ".git")
}

// Check the path for things that make working easier
func (w *PrDiscovery) checkForTools() {
	_, err := exec.LookPath("git-flow")
	if nil != err {
		w.hasGitFlow = false
	} else {
		w.hasGitFlow = true
	}
}

// Yeah, I know I can do `git branch` and sed the asterisk away. That's boring.
func (w *PrDiscovery) discoverCurrentBranch() {
	for _, command := range commandsToGetCurrentBranch {
		action := exec.Command(command[0], command[1:]...)
		result, err := action.CombinedOutput()
		if nil != err {
			continue
		}
		w.assumedCurrentBranch = strings.TrimSpace(string(result))
		break
	}

}

// With GitFlow, if the branch is currently active (or you broke it, which I do
// regularly), you can find its base in the config file. If you're not using
// GitFlow or it's been cleared, you can find the base branch on your own time.
// http://bbs.bugcode.cn/t/7634
func (w *PrDiscovery) discoverBase() {
	if w.hasGitFlow {
		action := exec.Command(
			"git",
			[]string{
				"config",
				"--local",
				fmt.Sprintf(
					"gitflow.branch.%s.base",
					w.assumedCurrentBranch,
				),
			}...,
		)
		result, err := action.CombinedOutput()
		if nil == err {
			w.suggestedBase = strings.TrimSpace(string(result))
		}
	}
}

// This convenience function builds a new PrDiscovery instance and runs through
// all its methods to generate suggestions for the PR payload
func CompletePrDiscovery() *PrDiscovery {
	discovery := PrDiscovery{}
	discovery.checkForTools()
	discovery.discoverGitConfigDirectory()
	discovery.discoverCurrentBranch()
	discovery.discoverBase()
	return &discovery
}

// This pulls a list of all the contributors on a project to use as a simple
// suggestion for adding reviewers
func (r *ReviewDiscovery) discoverLocalShortList() {
	action := exec.Command("git", []string{"shortlog", "-s", "-n"}...)
	result, err := action.CombinedOutput()
	if nil == err {
		r.localSuggestedReviewers = strings.Split(string(result), "\n")
	}
}

// This pulls a list of all the local contributors on a project to use as a
// simple suggestion for adding reviewers
func (r *ReviewDiscovery) discoverRemoteShortlist(owner, repo string) {
	body := getCollaboratorList(owner, repo)
	var collaborators map[string]interface{}
	_ = json.Unmarshal(body, &collaborators)
	fmt.Println(collaborators)
}

// This pulls a list of all the remote contributors on a project to use as a
// simple suggestion for adding reviewers
func CompleteReviewDiscovery(owner, repo string) []string {
	discovery := ReviewDiscovery{}
	discovery.discoverLocalShortList()
	discovery.discoverRemoteShortlist(owner, repo)
	return []string{}
}
