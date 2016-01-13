package repodb

// Build is a build of a repository.
type Build struct {
	Name  string
	Build string
	Lang  string

	// Jsonable structures.
	Struct interface{}
	Files  map[string]interface{}
}
