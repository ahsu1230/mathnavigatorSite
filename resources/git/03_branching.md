# Working with Git Branches

## Review: What is a Git Branch?

https://github.com/ahsu1230/StudentSandbox/blob/master/guide/02_branches.md

Every codebase has a *master* branch, which is the most up-to-date draft of a codebase. The codebase can also have *feature branches* which are like "alternative drafts" to the codebase which developers use to implement a feature without affecting each other.

Branches can also be local and remote. Local branches live on your computer while remote branches live on Github. So for example, the remote *master* branch is the most up-to-date codebase that lives on Github. A local *master* branch is supposed to be a copy of the remote *master* branch that lives on your computer. However, the local branch can often be out-of-sync, so developers need to manually update (using `git fetch`) their master branch to reflect the newest changes to the remote master.

## Using branches in a development team

When developing in a team, collaborators are often working on different features simultaneously. Sometimes these features are related, sometimes they are not, and sometimes they are unrelated, but edit the same files.

Branching is used to keep developers from affecting each other's work as much as possible. When developing a feature, every developer creates their own local feature branch, which is a copy of the remote master branch. The developer then works on this feature branch, usually adding many commits on top of this feature branch. As they work on this feature branch, the master branch is unaffected because only the copies are being worked on.

When the developer is ready, they may push their local feature branch to Github (remote). Once the feature branch is online, they can create a pull request for teammates to see the new feature code and suggest changes. The owner developer can commit more changes and push those new commits to the feature branch. The original pull request will be updated.

When the feature branch is finally approved, the feature branch can merge with the original remote master branch. At this point, the developer's feature is now available to the rest of the team-members when they work on their next feature!

## Suggested Workflow

In your computer Terminal, you can create a new local branch from master

```
git fetch
git checkout -b NEW_BRANCH origin/master
```

`git fetch` updates the remote so we can sync with Github to retrieve the latest updates to the repository. `git checkout ...` creates a new branch called NEW_BRANCH, and bases it off `origin/master` which refers to the remote *master* branch on Github. Please make sure to include the last part (`origin/master`). Otherwise, you end up creating a new branch off your current local branch (which could have changes from a previous feature).

This new local branch is created and should reflect what is currently on Github's master branch. From here, you can add commits as normal.

```
git status or git diff          // to view changed files or changed lines in files
git add -A
git commit -m "..."
```

When you are ready to create a pull request, use this command:

```
git push origin NEW_BRANCH
```

This will create a new remote branch called NEW_BRANCH on Github, which you can then use to create a Pull Request. If the branch already exists on Github, it will be "revived" from its stale state and allow you to create a Pull Request.

## Switching back

Let's say after you create a Pull Request and are waiting for a review, you want to work on another feature. If the feature is new, use the above `git checkout -b ...` command we mentioned earlier to create a new branch off the remote master. If the feature is still on a local branch, you can simply use `git checkout OLD_BRANCH_NAME` to switch back to the other branch.

**Note!** You cannot have any "pending" uncommitted changes when switching branches. All changes must be commited, stashed, or deleted before switching branches.

To delete your current all uncommitted changes, use `git reset --hard`. This will reset your branch to the last commit you've made. Any changes to files that were not committed will be lost.

## Pulling from Master

If you've been working on a feature for a while or the master branch has been progressing fast, Git may not allow you to push changes to the feature branch until you've pulled recent changes from master.

To pull from master, use:
```
git fetch
git pull origin master
```

This will create a commit that takes new changes from the remote master branch and updates your current feature branch. Once this is done, you may be able to push using `git push origin master`.

Sometimes, however, you may get a merge conflict! To resolve merge conflicts, see this [guide](./04_merge_conflicts.md).