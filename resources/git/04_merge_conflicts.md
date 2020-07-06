# What is a Merge Conflict?

- Watch the first minute of <https://www.youtube.com/watch?v=MzpW-k66XE8>

A Merge Conflict occurs when two branches are trying to merge but have conflicting changed lines/files. Usually, there are particular line(s) in a certain file that has been edited by both branches and Git is unable to automatically resolve how to combine these changes. When merge conflicts are detected, Git will automatically generate un-compileable lines into your code and prevent you from merging branches. Merge conflicts sound intimidating but they follow a pattern. Once you figure out that pattern, they're not too bad to work with!

## What do they look like?

Merge conflicts are often seen on Github Pull Requests. GitHub will show which files are unable to correctly merge and prevent the current feature branch from merging with the master branch.

![GIT_PR_MERGE_CONFLICT](../images/git_pr_merge_conflict.png)

Sometimes, they may occur when a developer calls `git pull origin master`. If the remote *master* branch has had many changes, the local feature branch may be too "out-of-sync" with the original *master* codebase so a pull is required. When the pull happens, there could be some files that the developer was working on but now have changed because of the new updates to *master*. You might see something like this:

```
$ git pull origin master
remote: Enumerating objects: 3, done.
remote: Counting objects: 100% (3/3), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0

Auto-merging ________________
Auto-merging ________________
CONFLICT (content): Merge conflict in README.md
Automatic merge failed: fix conflicts and then commit the result.
```

From here, use `git status` to view the list of files that must go through merge conflict resolution.

## How to resolve a merge conflict?

- Video: <https://www.youtube.com/watch?v=g8BRcB9NLp4>

Once you've identified which files have merge conflicts, open them up in a code editor to view these files. It should be very clear where the merge conflicts occur. There is un-compilable code that forces a developer to address those changes.

Here is an example of a Golang struct called `Class`.
```
type Class struct {
	Id              uint       `json:"id"`
	CreatedAt       time.Time  `json:"-" db:"created_at"`
	UpdatedAt       time.Time  `json:"-" db:"updated_at"`
	DeletedAt       NullTime   `json:"-" db:"deleted_at"`
	PublishedAt     NullTime   `json:"publishedAt" db:"published_at"`
	ProgramId       string     `json:"programId" db:"program_id"`
	SemesterId      string     `json:"semesterId" db:"semester_id"`
	ClassKey        NullString `json:"classKey" db:"class_key"`
	ClassId         string     `json:"classId" db:"class_id"`
	LocationId      string     `json:"locationId" db:"location_id"`
	Times           string     `json:"times"`
<<<<<<< HEAD
	PaymentNotes    NullString `json:"paymentNotes" db:"payment_notes"`
-=======
	FullState       int        `json:"fullState" db:"full_state"`
	PricePerSession NullUint   `json:"pricePerSession" db:"price_per_session"`
>>>>>>> a711987b69fd2bf2cc26c2b1448318acd198725c
}
```

As you can see, Git is having trouble identifying what fields should be part of Class. There is code portion from two branches separated by `=======`. The top portion (after `HEAD`) are local changes (what you had before the `git pull`). The bottom portion is "incoming" changes (changes coming from the merging branch). So if you're pulling from master, the bottom code portion is referring to the most up-to-date changes in the remote master branch codebase.

Let's look at the changes in more detail. When resolving merge conflicts, you MUST keenly observe both code segments. On one hand, one branch (local branch, current changes) is saying `PaymentNotes` should be part of the struct. The other branch (incoming branch) is saying the `FullState` and `PricePerSession` should be part of the struct. This usually happens if two developers simultaneously made changes to this Class struct. 

Turns out that in this case, we actually want both to be part of the Class struct. So we can combine the two code segments together and remove the `<<<<<<`, `======`, `>>>>>>` that would fail the code compilation.

```
type Class struct {
	Id              uint       `json:"id"`
	CreatedAt       time.Time  `json:"-" db:"created_at"`
	UpdatedAt       time.Time  `json:"-" db:"updated_at"`
	DeletedAt       NullTime   `json:"-" db:"deleted_at"`
	PublishedAt     NullTime   `json:"publishedAt" db:"published_at"`
	ProgramId       string     `json:"programId" db:"program_id"`
	SemesterId      string     `json:"semesterId" db:"semester_id"`
	ClassKey        NullString `json:"classKey" db:"class_key"`
	ClassId         string     `json:"classId" db:"class_id"`
	LocationId      string     `json:"locationId" db:"location_id"`
	Times           string     `json:"times"`
	PaymentNotes    NullString `json:"paymentNotes" db:"payment_notes"`
	FullState       int        `json:"fullState" db:"full_state"`
	PricePerSession NullUint   `json:"pricePerSession" db:"price_per_session"`
}
```

Fortunately, this merge conflict example is very easy to resolve. But the process of more complicated merge conflicts is the same.

- Identify which files have merge conflicts
- Observe which lines of the codebase have incompatible changes (and observe the contents of both codeblocks)
- Select the final change that needs to persist to the *master* branch.
- Remove extraneous characters from the merge conflict (`<<<`, `===`, `>>>`).
- Save and commit.

Once we are done resolving, input this in Terminal.

```
git status          <- checks for any other remaining files with merge conflicts
git add -A
git commit          <- Creates a commit from the merge process
```

`git commit` will take you to the Vim COMMAND Mode which looks like this:

Use Esc -> `:wq` to exit Vim and save the merge commit.
When you are finished, you can push the commit to the Pull Request. You should now see that the merge conflicts have been resolved and you should be able to merge the branch!