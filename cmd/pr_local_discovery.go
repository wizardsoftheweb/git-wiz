package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// Created the same thing to hold the discovery work
// I should probably separate this out by provider but I'm not there yet
type WotwPrRequest struct {
	// Composition
	GitHubPrRequest
	// Theoretically the branch we're on right now
	suggestedBase string
	// Theoretically tracked in git-flow
	suggestedHead string
	// A list of people that might be good to ask to review the PR
	suggestedReviewers []string
	// Theoretically the title of this branch
	suggestedTitle string
	// If it's there, this becomes much easier.
	hasGitFlow bool
	// This comes from the CLI before the app reaches out to the user so it
	// could be suspect.
	assumedCurrentBranch string
	// The root path of the repo that contains what I need.
	gitConfigDirectory string
}

var (
	// Hopefully one of these will actually work
	commandsToGetCurrentBranch = [][]string{
		{"git", "branch", "--show-current"},
		{"git", "rev-parse", "--abbrev-ref", "HEAD"},
		{"git", "symbolic-ref", "--short", "HEAD"},
	}
)

var Demo WotwPrRequest

// If we're in the repo, rev-parse should get us there. If not, we'll have to
// prompt for it.
func (w *WotwPrRequest) discoverGitConfigDirectory() {
	action := exec.Command("git", []string{"rev-parse", "--show-toplevel"}...)
	result, err := action.CombinedOutput()
	if nil != err {
		return
	}
	w.gitConfigDirectory = filepath.Join(string(result), ".git")
}

// Check the path for things that make working easier
func (w *WotwPrRequest) checkForTools() {
	_, err := exec.LookPath("git-flow")
	if nil != err {
		w.hasGitFlow = false
	}
	w.hasGitFlow = true
}

// Yeah, I know I can do `git branch` and sed the asterisk away. That's boring.
func (w *WotwPrRequest) discoverCurrentBranch() {
	for _, command := range commandsToGetCurrentBranch {
		action := exec.Command(command[0], command[1:]...)
		result, err := action.CombinedOutput()
		if nil != err {
			continue
		}
		w.assumedCurrentBranch = string(result)
		break
	}
}

// With GitFlow, if the branch is currently active (or you broke it, which I do
// regularly), you can find its base in the config file. If you're not using
// GitFlow or it's been cleared, you can find the base branch on your own time.
// http://bbs.bugcode.cn/t/7634
func (w *WotwPrRequest) discoverBaseBranch() {
	if w.hasGitFlow {
		action := exec.Command(
			"git",
			[]string{
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
			if "" != w.suggestedBase {
				return
			}
		}
	}
}

// This pulls a list of all the contributors on a project to use as a simple
// suggestion for adding reviewers
func (w *WotwPrRequest) discoverShortList() {
	action := exec.Command("git", []string{"shortlog", "-s", "-n"}...)
	result, err := action.CombinedOutput()
	if nil == err {
		w.suggestedReviewers = strings.Split(string(result), "\n")
	}
}
