# Hooks and Filters

This directory contains (will contain) some things I've been playing with off and for some time. I don't currently have a unified hook setup but I will one of these days.

Here's a list of all the hooks I know and might potentially one day use:
* [Vanilla githooks](https://git-scm.com/docs/githooks#_hooks)
* [`gitflow-avh` hooks](https://github.com/petervanderdoes/gitflow-avh/wiki/Reference:-Hooks-and-Filters) (probably won't work with vanilla but I don't know)

## Prereqs

### XDG Base Directory Specification

Skim [the spec](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) really fast. While you don't need everything set up now, I'd recommend you take the time to set up `$XDG_CONFIG_HOME`.

```
mkdir -p $HOME/.config/git
echo 'export XDG_CONFIG_HOME=$HOME/.config' >> .whateverrc  
```

`git` uses the XDG dirs for config. This is useful because it cuts down on the pollution of your `$HOME` and better organizes files. It's not 1989 anymore; dotfiles below in neat, organized subdirectories, not thrown all together in a giant mess.

### `GIT_DIR` New and Old

If you do any research (including [reading the docs](https://git-scm.com/book/en/v2/Git-Internals-Environment-Variables)), you'll see a `GIT_DIR` variable talked about. It references to the location of the repo's `.git` directory. I'm not entirely sure about the behavior of this environment variable (never have been) and in recent years I've heard more chatter about it be being set inconsistently if at all. If you work in an environment when it work for you, rad. I have switched to this command because it's portable and I don't have to research if new environments can handle it.

```shell-session
$ git rev-parse --show-toplevel
/home/rick.james/some-repo
$ echo "$(git rev-parse --show-toplevel)/.git"
/home/rick.james/some-repo/.git
```

## Automatic Installation and Organization

I haven't built a (complete, usable) solution yet but plenty of others have. There are some pretty neat things out there! 

## Manual Installation and Organization

There are three approaches to manual installation and organization that I see. There might be more and I'd love to know if there are. You can

* manually choose when to update it by keeping your hooks in some directory not referenced by active config,
* store them all in the same place, a template directory, that will automatically install them with new repos, or
* store them all in the same place and have every repo use the same hooks directory.

My (eventual) solution is an automatable and programmatic solution but I'm not there yet today.

### Always Manually Install

I'd highly recommend carving yourself out a directory in `$XDG_CONFIG_HOME/git` for this. Even if you aren't going to apply them globally, it sucks to lose or forget them over time.

Setting this up looks something like this:

```shell-session
mkdir -p $XDG_CONFIG_HOME/git/all-hooks
touch $XDG_CONFIG_HOME/git/all-hooks/this-is-fake
mkdir my-new-repo
cd my-new-repo
git init
# Copying means you have to manually update it
cp $XDG_CONFIG_HOME/git/all-hooks/this-is-fake .git/hooks
# Symlinking means less updating
# I personally prefer to (mostly) use resolved, absolute refs for symlinks
ln -sf /home/rick.james/.config/git/all-hooks/this-is-fake $(git rev-parse --show-toplevel)/.git/hooks
```

#### Pros

You to choose what goes where and when it's updated (unless you symlink)

#### Cons

You have to keep track of what goes where and when it's (unless you symlink).

### Install With `git init`

If you get tired of manually linked those files, you can always install the hooks as part of `git init`. The command copies [the contents of the template directory](https://git-scm.com/docs/git-init#_template_directory) into the `git` directory created by `git init`. Rerunning the command will copy new or updates files into the existing directory. This is what mine looks like:

```shell-session
$ git --version
git version 2.23.0.37.g745f681289
$ tree /usr/share/git-core/templates
/usr/share/git-core/templates
├── branches
├── description
├── hooks
│   ├── applypatch-msg.sample
│   ├── commit-msg.sample
│   ├── fsmonitor-watchman.sample
│   ├── post-update.sample
│   ├── pre-applypatch.sample
│   ├── pre-commit.sample
│   ├── prepare-commit-msg.sample
│   ├── pre-push.sample
│   ├── pre-rebase.sample
│   ├── pre-receive.sample
│   └── update.sample
└── info
    └── exclude

3 directories, 13 files
```

This is what happens when I `init`:
```shell-session
$ cd some-new-directory
$ tree -a .
.

0 directories, 0 files
$ git init
Initialized empty Git repository in /home/rick.james/some-new-directory/.git/
$ tree -a .
.
└── .git
    ├── branches
    ├── config
    ├── description
    ├── HEAD
    ├── hooks
    │   ├── applypatch-msg.sample
    │   ├── commit-msg.sample
    │   ├── fsmonitor-watchman.sample
    │   ├── post-update.sample
    │   ├── pre-applypatch.sample
    │   ├── pre-commit.sample
    │   ├── prepare-commit-msg.sample
    │   ├── pre-push.sample
    │   ├── pre-rebase.sample
    │   ├── pre-receive.sample
    │   └── update.sample
    ├── info
    │   └── exclude
    ├── objects
    │   ├── info
    │   └── pack
    └── refs
        ├── heads
        └── tags

10 directories, 15 files
```
If you didn't know, you've had hook templates sitting in, most likely, every single repo you've ever made.

The template directory is the first item found from this list:

* the `--template` flag
* the `$GIT_TEMPLATE_DIR` env variable
* the `init.templatedir` config variable (which could either be `global` or `system`)
* the `/usr/share/git-core/templates` directly, which come with `git.

#### Pros

You get everything everywhere and can specifically not update the hooks in repos where you don't want to see new versions.

#### Cons

You're stuck with everything everywhere and you have to manually `git init` (or copy) files every time you want an update.

### Use a Global Directory

You can configure `git` to source source its hooks from a single location [using the `core.hookspath` option](https://git-scm.com/docs/git-config#Documentation/git-config.txt-corehooksPath). This is basically identical to the previous one except you run the same version the script everywhere. When you update, everything gets it.

#### Pros

Everything is in a single location and always up-to-date.

#### Cons

Everything runs the same version whether you like it or not and you have to make sure everything everywhere works before updating.  

## Hooks Here So Far

### Prune Branches That Have Been Merged (Like Old Prefix Branches)

`git flow <prefix> finish` does this automatically for you but it's a solo action. If you share code responsibility with anyone else, you should always PR your code (even if it doesn't get looked at) to build solid habits of openness and transparency. Plus good reviews make things better. That being said, they also usually lead to other people merging your branch and, even though it's horrible to say, not immediately deleting the now useless prefix branch from the remote.

Your first instinct might be to `git {fetch,pull} --prune --all`, but those commands only update the references.  They won't delete stale ones. AFAIK no tool will do that automatically. There are semi-legitimate concerns about losing code that tend to go away the more your team becomes educated about `git` that usually prevent this process from getting started.

Our goal is remove dead weight. A branch is consider dead weight if it's been merged into its base, eg the `dev` branch for `feature` branches. Once prefix branch (ie anything that isn't `dev`, `master`, or one of the fancier `support`/`legacy`/`whatever` trunks) has been merged, it shouldn't be touched again. It should go away. Its changes are in its root, which could also have changes from other sources, so we should start new work from there.

There's a very convenient way to list branches that have been merged with our work. I say "our work" because, in a team situation, we need to be very cognizant that our automation solutions can affect others. While there might be another way to go about this, choosing branches that have been merged into ours should ruffle the least amount of feathers. Branches are just specific commits that end a chain of work. If a branch has been merged into ours, its tip must be a descendant which means its contents have to be behind us. In other words, we're not dropping someone else's changes. That's now how merged branches work.


```shell-session
$ git branch --all --merged
  dev
* feature/subfeature-branch
  feature/root-branch
  feature/old-deletable-feature
  master
  remotes/origin/dev
  remotes/origin/feature/subfeature-branch
  remotes/origin/feature/root-branch
  remotes/origin/feature/old-deletable-feature
  remotes/origin/master
  remotes/other-remote/dev
  remotes/other-remote/feature/subfeature-branch
  remotes/other-remote/feature/root-branch
  remotes/other-remote/feature/old-deletable-feature
  remotes/other-remote/master
  remotes/third/remote/time/dev
  remotes/third/remote/time/feature/subfeature-branch
  remotes/third/remote/time/feature/root-branch
  remotes/third/remote/time/feature/old-deletable-feature
  remotes/third/remote/time/master
```

This shows that the current `dev` and `master` branches, both local and remote, have been merged into my active branch. There are also some other branches, local and remote, that have been merged but not deleted. We don't want to delete the entire list; we want to only delete the cruft. This means we have to be very careful with what we chose to delete, especially on the remote(s).


Since we're following `gitflow` we know everything that's been merged into its `dev` branch (other than `master`) can be deleted. Anything newer than the `dev` branch might be an active feature touched by someone else on the team. Locally it could be another feature that's been tabled for whatever reason.

Locally that's very easy.

```shell-session
$ git branch --merged=dev
  dev
  feature/old-deletable-feature
  master
$ git branch --merged=dev | grep -vE 'dev|master'
  feature/old-deletable-feature
$ git branch --merged=dev | grep -vE 'dev|master' | xargs -n 1 git branch -d 
Deleted branch feature/old-deletable-feature (was 8a3b292).
```

With a single remote, it's still fairly easy but needs more steps.

```shell-session
$ git branch --remote --merged=dev \
    | grep origin \
    | grep -vE 'origin/(dev|master)$' \
    | sed 's@origin/@origin @' \
    | xargs -n 2 bash -c 'echo git push $0  --delete $1'
git push origin --delete feature/another-deletable-feature
```
With multiple remotes, it seems harder but just requires some some judicious variable creation before piping. Since remotes can be named like branches or refs (ie like folders) we can't make assumption about their final naming.

```shell-session
$ git remote \
    | xargs -n 1 bash -c '\
        REMOTE=$0; \
        BRANCHES=$( \
            git branch --remote --merged="$REMOTE"/dev \
                | grep $REMOTE \
                | sed "s@$REMOTE/@@" \
                | grep -vE " (master|dev)$" \
        ); \
        [ -z "${BRANCHES[@]}" ] || echo ${BRANCHES[@]} \
            | xargs -n 1 bash -c  "echo git push $REMOTE --delete "'
git push origin --delete feature/another-deletable-feature
git push other-remote --delete feature/another-deletable-feature
git push third/remote/time --delete feature/another-deletable-feature
```

```shell-session
$ git branch --all --merged
  dev
* feature/subfeature-branch
  feature/root-branch
  master
  remotes/origin/dev
  remotes/origin/feature/subfeature-branch
  remotes/origin/feature/root-branch
  remotes/origin/master
  remotes/other-remote/dev
  remotes/other-remote/feature/subfeature-branch
  remotes/other-remote/feature/root-branch
  remotes/other-remote/master
  remotes/third/remote/time/dev
  remotes/third/remote/time/feature/subfeature-branch
  remotes/third/remote/time/feature/root-branch
  remotes/third/remote/time/master
```

Note that, in order to run the many-remote commands, you'll need to remove that last `echo` so the command is executed instead of printed.

After all that was said and done, we still have a fairly large number of branches live. This can't be managed programmatically (unless you set some rules up first). In each setting, `dev`, `master`, and `feature/root-branch` are in `feature/subfeature-branch`. We need `dev` and `master` as the sources of truth. The whole point of running a subfeature is merge it back in, so we can't nuke `feature/subfeature-branch`. If we didn't have those requirements, we could manually delete things (like who needs two separate mirrors of `origin`?).


