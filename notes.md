- let's build an online file system.
- navigation bar on top

< branch @ version/time > | user(repo) > package > file

first version:
- one file system, only save snapshots
- can store only changed files (each file about 300 lines anyways)
- each snapshot is logically a dict of filenames to file contents

actions:
- add/set/delete file in a snapshot
- all actions to a snapshot are serialized
- merge snapshot must: no file conflict, compiles, tests passing

session is a user name
- NewCommit(base int) int
- ListFiles(c int) []string
- Read(c int, name string) *string
- Write(c int, name string, s *string) error
- Submit(c int) // mark as read only now
- Merge(base, other int) (int, error) // this is the hardest one
- NewBranch(name string, c int)
- BranchMerge(name string, c int) error

each user has its own repo
each repo has a list of snapshots
a branch is a list of snapshots
the master is a special branch
the repo owner can merge a branch into master
the branch owner can merge any other branch into its branch
only the master branch is importable from other 

- auto saving snapshots
- snapshots can have names
- online meld like diff support (optional)
- all files are text files
- folders work like a file explorer/browser
- command line support in a folder (optional)

backend
- for simplicity, all operations are logged and serialized.
- each operation will evolve into a new version.
- one can clone from one tag into another tag.
- a branch is hence fundamentally a cl, or a diff
- a branch has a base (the diff base)
- each branch snapshot has a father
- so, we save branches always in diffs, with file/folder add/delete/changes.
- and we save the master thread always in snapshots, with deduped files.

