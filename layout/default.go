package layout

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
)

const (
	errorUnknownProject = "Could not infer project type"
)

// AppFs is a interface to the filesystem, used for mocking tests
var AppFs = afero.NewOsFs()

// DefaultProject represents default assumptions gb makes while
// building projects. Gb entire idea is based on Convention-over-Configuration,
// and Gb infers as much as possible from project folder sturcture. Similar
// to Ruby or Java's Maven.
type DefaultProject struct{}

func (DefaultProject) Get(projectRoot string, buildRoot string) (ProjectInfo, error) {
	var meta ProjectInfo

	meta.Path.Includes = filepath.Join(projectRoot, "include")
	meta.PublicIncludes = meta.Path.Includes
	meta.Path.Sources = filepath.Join(projectRoot, "src")
	meta.Path.Tests = filepath.Join(projectRoot, "test")

	includeOk, _ := afero.DirExists(AppFs, meta.Path.Includes)
	srcOk, _ := afero.DirExists(AppFs, meta.Path.Sources)
	testOk, _ := afero.DirExists(AppFs, meta.Path.Tests)

	if includeOk {
		meta.HasPublicIncludes = true
		if srcOk {
			meta.Type = Library
		} else {
			meta.Type = HeaderOnly
		}
	} else if srcOk {
		meta.HasPublicIncludes = false
		meta.Type = Application
	} else {
		meta.Type = Unknown
		return meta, fmt.Errorf(errorUnknownProject)
	}

	if srcOk {
		sources, err := GetProjectFiles(filepath.Join(projectRoot, "src"), "**/*.cpp", buildRoot)
		if err != nil {
			return meta, err
		}
		meta.Sources = sources
	}

	if testOk {
		meta.HasTests = true

		tests, err := GetProjectFiles(filepath.Join(projectRoot, "test"), "**/*.cpp", buildRoot)
		if err != nil {
			return meta, err
		}
		meta.Tests = tests
	}

	return meta, nil
}

// NewDefaultProjectLayout returns implementation of ProjectLayout
// for Gb default layout
func NewDefaultProjectLayout() ProjectLayout {
	return DefaultProject{}
}
