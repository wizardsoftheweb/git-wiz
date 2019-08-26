package git_credentials_store

import (
	"os/exec"
	"regexp"
	"strconv"
)

var DoPathsMatter = PathConfigMonitor

func PathConfigMonitor() bool {
	blob := CollectGitConfig()
	return CheckPathValue(blob)
}

// var PatternGitConfigLine = regexp.MustCompile(`(?mi)^credential\.(.*?)\.*(helper|usehttppath|username)\s*(.*)\s*$`)
var PatternConfigUseHttpPath = regexp.MustCompile(`(?mi)^credential\.(.*?)\.*(usehttppath)\s*(.*)\s*$`)

func CollectGitConfig() string {
	commandArgs := []string{"config", "--get-regexp", "credential*"}
	result, _ := exec.Command("git", commandArgs...).CombinedOutput()
	return string(result)
}

func CheckPathValue(blob string) bool {
	pathsMatter := false
	matches := PatternConfigUseHttpPath.FindAllStringSubmatch(blob, -1)
	for _, match := range matches {
		if "" != match[3] {
			pathsMatter, _ = strconv.ParseBool(match[3])
		}
	}
	return pathsMatter
}
