package fs8

// FileSys is a general versioning file system interface.
type FileSys interface {
	Login(user string) Session
}

// Commit is a file system commit.
type Commit struct {
	ID      int
	Message string
}

// Session is a file system login session.
type Session interface {
	CheckOut(base int) (int, error)
	Commits(n int) ([]*Commit, error)
	ListFiles(c int) ([]string, error)
	Read(c int, name string) (*string, error)
	Write(c int, name string, s *string) error
	Submit(c int, message string) error
	Merge(base, other int) (int, error)
	NewBranch(name string, c int) error
	Accept(branch string, c int) error
}
