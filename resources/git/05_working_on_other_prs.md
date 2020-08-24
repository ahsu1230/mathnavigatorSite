# Working on Other People's Pull Requests

Being able to work on another person's PR can come in handy when minor fixes need to be applied or when you're working on a feature with another person. Working on another person's PR is fairly simple:

1. Fetch to get the latest changes from Github

```
git fetch
```

2. Checkout the PR's branch and get its latest changes

```
git checkout OTHER_BRANCH
git pull origin OTHER_BRANCH
```
You may run into merge conflicts at this step - refer to this [guide](./04_merge_conflicts.md) to resolve them.

3. Make your desired changes
4. Commit and push

```
git add -A
git commit -m "Made changes"
git push origin OTHER_BRANCH
```

If you get an error about the current branch being behind, you haven't pulled the latest changes. Run `git pull origin OTHER_BRANCH` again and resolve merge conflicts if any.
