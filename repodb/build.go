package repodb

// Build is a build of a repository.
type Build struct {
	Name  string
	Build string
	Lang  string

	Struct interface{}
	Files  map[string]interface{}
}

// LatestBuild is a structure for saving the latest build of a repo.
type LatestBuild struct {
	Repo   string
	Build  string
	Lang   string
	Struct []byte
}