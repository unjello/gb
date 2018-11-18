package layout

const (
	// Application describes project that results in executable
	Application = iota
	// Library is a project that other can depend on (include)
	Library = iota
	// HeaderOnly is a library that does not require linking
	HeaderOnly = iota
	// Unknown means project could not be identified
	Unknwon = iota
)

// ProjectLayout provides set of functions for describing C++ project
type ProjectLayout interface {
	Get(root string) (ProjectInfo, error)
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
	Type           int8
	Name           string
	PublicIncludes string
	Sources        []SourceFile
	Tests          []SourceFile
	TestsIncludes  []string
}
