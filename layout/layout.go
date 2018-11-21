package layout

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
	"github.com/evilsocket/islazy/tui"
)

const (
	// Application describes project that results in executable
	Application = iota
	// Library is a project that other can depend on (include)
	Library = iota
	// HeaderOnly is a library that does not require linking
	HeaderOnly = iota
	// Unknown means project could not be identified
	Unknown = iota
)

// ProjectLayout provides set of functions for describing C++ project
type ProjectLayout interface {
	Get(projectRoot string, buildRoot string) (ProjectInfo, error)
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
	Type              int8
	Name              string
	PublicIncludes    string
	Sources           []SourceFile
	Tests             []SourceFile
	TestsIncludes     []string
	HasTests          bool
	HasPublicIncludes bool
	Path              struct {
		Includes string
		Sources  string
		Tests    string
	}
}

func GetProjectFiles(root string, globPattern string, buildRoot string) ([]SourceFile, error) {
	sourceGlob := filepath.Join(root, globPattern)
	files, err := doublestar.Glob(sourceGlob)
	if err != nil {
		log.Fatal("Failed to find sources using pattern: " + tui.Red(sourceGlob))
		log.Fatal(err.Error())
		return nil, err
	}

	sourceFiles := make([]SourceFile, 0)
	for _, file := range files {
		relPath, err := filepath.Rel(buildRoot, file)
		if err != nil {
			log.Fatal("Failed to reach " + tui.Dim(file) + " from build dir: " + tui.Dim(buildRoot))
			return nil, err
		}
		base := filepath.Base(file)
		ext := filepath.Ext(base)
		baseName := strings.TrimRight(base, ext)
		sourceFiles = append(sourceFiles, SourceFile{file, relPath, baseName, ext})
	}
	return sourceFiles, nil
}
