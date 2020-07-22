# Common Git Commands

This is a short list of Git commands you will frequently use when working on Git projects. This list is meant more of a review of common commands. If you need a refresher of concepts, check out this resource as well <http://rogerdudler.github.io/git-guide/>.

## General Workflow

For simplicity sake, your overall workflow will probably consist of this series of commands.

1. Update remotes
2. Checkout a branch (local or remote)
3. Pull from Remote (if necessary)
4. *Make code changes...*
5. Check Status / Diffs
6. Add Files
7. Commit change
8. Push commits (if want commits on remote)
9. Repeat steps 4-8 until finished with feature
10. Create a Pull Request on Github
11. Repeat steps 4-8 again if any additional changes are necessary
12. Merge when code-owner approves pull request

***TIP*** You CANNOT switch branches when you have any un-committed local changes. If you need to switch branches in the middle of working on a feature, you should first add and commit all changed files, undo all of the changes, or stash them.

***TIP*** Always be sure to update remotes as often as possible. A teammate might have updated file(s) you were working on - you probably want the latest updates! Update remotes ASAP so you can save yourself the trouble of dealing with merge conflicts later!

## Managing branches

```git
git branch
git checkout BRANCH_NAME
git checkout -b NEW_BRANCH_NAME
git checkout -b NEW_BRANCH_NAME origin/master
git branch -D BRANCH_NAME
```

1. List all local branches
2. Make an existing branch your current branch
3. Creates a new branch (off of current branch) and makes it your current branch
4. Creates a new branch (off of remote master) and makes it your current branch
5. Deletes a local branch

For more details about managing branching in your workflow, visit [here](./03_branching.md).

## Managing Changed Files

### Checking Status

```git
git status
git diff
```

1. View a list of changed files
2. View content of all changed files

### Staging Files

```git
git add FILE_NAME
git add -A
```

1. Stage a single file (not committed yet)
2. Stage all changed files (not committed yet)

### Unstaging Files

```git
git checkout FILE_NAME
git reset FILE_NAME
git reset --hard
```

1. Undo all changes to a single file
2. Unstage a single file
3. **CAUTION** Unstage and undo all changed files. Reset current branch to before you made any changes.

## Managing Commits

```git
git commit -m "message"
git commit --amend

git revert COMMIT_HASH
git log
```

1. Commit all "staged" files with a message
2. Update the message of last commit
3. Creates a commit that undos the last commit
4. View a history of all latest commits

## Push and Pull

```git
git fetch
git pull origin master
git pull origin LOCAL_BRANCH_NAME
```

*TIP* It's advised to always call `git fetch` before using `git pull`. That way, you can download the latest updates from the remote repository. Without doing this, your remote may be out-of-date and end up pulling old updates from the remote repository. This may cause merge conflicts later.

1. Update your current references to point to the latest changes in remote repository
2. Pull the latest changes from remote `master` to your current local branch. You will most likely be using this pull command.
3. Pull the latest changes from the remote branch to your local branch. You probably only need this when multiple users are writing to your remote branch (i.e. another teammate, or if you made edits using the Github website).

## Stash

```git
git stash list
git stash
git stash save -m ""
git stash show
git stash apply
git stash --help
```

1. List all stashes on computer
2. Stash all staged files (message by default will just be a random hash)
3. Stash all staged files with a message
4. Show the most recent stash
5. Apply the most recent stash and remove it from the stash list (like a pop function)
6. See more details about stash command

## Managing a Repository

```git
git init
git clone
```

1. In current directory, create a new Git repository!
2. Checkout a remote Git repository. Use this to download someone else's code!

## Managing Remotes

```git
git fetch
git remote --help
```

1. Updates all remotes
2. View more remote commands


## Ignore Files

To ignore files, add your file(s) to the `.gitignore` file inside the root directory of the git repository. Note that it is a "hidden file" so look for this file using Command.
