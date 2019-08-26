package git_credentials_store

import (
	"regexp"
)

const testContent = `\
protocol=https
password=qqq
path=seven
host=rick.james
username=bobby
golang=something
`

var PatternSingleInputLine = regexp.MustCompile(`(?m)^\s*(protocol|username|host|path)\s*=\s*(.*)\s*$`)

func ParseSearchInput(input string) map[string]string {
	matches := PatternSingleInputLine.FindAllStringSubmatch(input, -1)
	incoming := make(map[string]string)
	for _, value := range matches {
		incoming[value[1]] = value[2]
	}
	return incoming
}
