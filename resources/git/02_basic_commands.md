## **Basic Git Commands to use:**

  - **Creating and Switching branches:**

Creates a new branch of the name you input
```
git branch --create NEW_BRANCH_NAME
```
Lets you see how many branches there are
```
git branch
```
Switches from your current branch to the desired branch
```
git checkout BRANCH_NAME
```
  - **Status, Add, Commit and Push**

"status" - Lets you see which changes have been staged, which haven't, and which files aren't being tracked
```
git status
```
"add" - Tells Git that you want to include updates to a particular file in the next commit
```
git add FILE_NAME
```
"commit" - Creates a new commit containing the current contents of the index and the given log message describing the changes
```
git commit -m "My First Commit Message"
```
"push" - Used to upload local repository content to a remote repository.
```
git push origin NAME_OF_BRANCH
```
  - **Resetting & Reverting**
https://www.atlassian.com/git/tutorials/resetting-checking-out-and-reverting

  - **Other Commands that are Useful**
"help <verb>" or < "<verb> --help" - a list of the most commonly used git commands are printed on the standard output
```
git <verb> --help
      or
git help <verb>
```

## Creating a Pull Request
Once your changes are pushed, come to this Github repo website and look for the Branches tab. You should see your branch and there should be a "Create Pull Request" button next to your branch.
Once a Pull Request is created, tag me using @ahsu1230 to notify me that this pull request has been created.
When I give the go ahead, merge your changes and now your code will be part of this repo!
