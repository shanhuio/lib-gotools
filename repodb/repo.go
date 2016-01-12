package repodb

// RepoBuild is a build of a repository.
type RepoBuild struct {
	Name  string
	Build string
	Lang  string

	// Jsonable structures.
	Struct interface{}
	Files  map[string]interface{}
}
