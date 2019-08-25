package main

import "github.com/wizardsoftheweb/git-wiz/cmd"

func main() {
	err := cmd.Execute()
	whereErrorsGoToDie(err)
}

func whereErrorsGoToDie(err error) {
	if nil != err {
		panic(err)
	}
}
