package layout

// ProjectLayout provides set of functions for describing C++ project
type ProjectLayout interface {
	IsLibrary(root string) (bool, error)
	IsApplication(root string) (bool, error)
}

// SourceFile handles metainformation for file that can be compiled
type SourceFile struct {
	FullPath  string
	RelPath   string
	BaseName  string
	Extension string
}

// ProjectInfo contains all information needed to build a project
type ProjectInfo struct {
	Name           string
	PublicIncludes string
	Sources        []SourceFile
	Tests          []SourceFile
	TestsIncludes  []string
}
