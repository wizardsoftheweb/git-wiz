# `git-wiz`

[![CircleCI](https://img.shields.io/circleci/build/github/wizardsoftheweb/git-wiz/dev)](https://circleci.com/gh/wizardsoftheweb/git-wiz/tree/dev)
[![Coverage Status](https://img.shields.io/coveralls/github/wizardsoftheweb/git-wiz/dev)](https://coveralls.io/github/wizardsoftheweb/git-wiz?branch=dev)
[![GoDoc](https://godoc.org/github.com/wizardsoftheweb/git-wiz?status.svg)](https://godoc.org/github.com/wizardsoftheweb/git-wiz)


I dunno; I just like building `git` tooling. This repo holds some experiments and (will eventually) exposes those far enough long via a binary that can be dropped in and discovered by `git`.

## Installing `git-wiz`

### With `go`

If you just want the binary,
```shell-session
go install ./... github.com/wizardsoftheweb/git-wiz
```

If you want to play with it in your code (haven't tried that, tbh),
```shell-session
go get -u ./... github.com/wizardsoftheweb/git-wiz
```

### Without `go`

Unless you have friends that use `git-wiz` that will also send you a compiled binary, I think it might be currently impossible to install `git-wiz` without `go` because I don't currently ship a binary. That's actually a really interesting question and I don't know enough about that side of Golang to say one way or the other. You'd certainly have to go through a lot of trouble figuring out how to compile and install it without `go`. That sounds like a challenge someone somewhere might find fun. I dunno. Coders are weird. 

## Tools currently in `git-wiz` even if they might be ready

Please note this list may be out of date; use the files here in the repo as the actual source of truth. Alternatively use the `git-wiz` binary itself; it might be in the repo but not yet compiled (although that's a super easy fix).

### `pr`

Creates a PR programmatically using information gleaned from the working directory and repo. I built it to speed up my collaborative `gitflow` process and will eventually either hook it in or duct tape it on. Right now it only supports GitHub but I do plan to add other providers eventually.

## Experiments not yet in `git-wiz`

Please note this list may be out of date; use the files here in the repo as the actual source of truth.

### `git-credential-crypt`

Fun with the `git` credential API
