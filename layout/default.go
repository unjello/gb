package layout

import (
	"fmt"
	"path/filepath"

	"github.com/evilsocket/islazy/log"
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

func (DefaultProject) Get(root string) (ProjectInfo, error) {
	var meta ProjectInfo

	isLib, _ := isLibrary(root)
	if isLib {
		meta.Type = Library
	} else {
		isApp, _ := isApplication(root)
		if isApp {
			meta.Type = Application
		} else {
			log.Error(errorUnknownProject)
			meta.Type = Unknwon
			return meta, fmt.Errorf(errorUnknownProject)
		}
	}

	return meta, nil
}

// IsLibrary returns true if a project has `include` folder,
// that has public headers for the project. Those are headers
// that can be included by other projects. Binary projects will
// not have anyone depend on them, therefore they will not expose
// any headers to noone.
// This function does not verify existance of `src` folder, as
// header-only libraries are as good as any. (if not better:)
func isLibrary(root string) (bool, error) {
	return includeFolderExists(root)
}

// IsApplication returns true if project has NOT `include` folder,
// and includes `src` folder. This does not verify whether there
// are source files inside, or if any of them actually has `main`
// function (or equivalnt). We lave it up to compiler, to decide
// whether to fail or note.
func isApplication(root string) (bool, error) {
	isLibrary, _ := isLibrary(root)
	if isLibrary {
		return false, nil
	}

	return srcFolderExists(root)
}

// NewDefaultProjectLayout returns implementation of ProjectLayout
// for Gb default layout
func NewDefaultProjectLayout() ProjectLayout {
	return DefaultProject{}
}

func srcFolderExists(root string) (bool, error) {
	srcPath := filepath.Join(root, "src")
	ok, err := afero.DirExists(AppFs, srcPath)
	if ok {
		return ok, err
	}

	return false, err
}

func includeFolderExists(root string) (bool, error) {
	includePath := filepath.Join(root, "include")
	ok, err := afero.DirExists(AppFs, includePath)
	if ok {
		return ok, err
	}

	return false, err
}
