package layout

import (
	"testing"

	"github.com/evilsocket/islazy/log"
	"github.com/spf13/afero"
)

func init() {
	log.Level = log.FATAL + 1
}

func TestIsLibraryWhenOnlyIncludeFolderPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"include", "build", "doc"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	ok, err := layout.IsLibrary("./")

	if !ok {
		t.Errorf("Expected true, but got %t: %q", ok, err)
	}
}

func TestIsLibraryWhenIncludeAndSrcFoldersPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"include", "build", "doc", "src"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	ok, err := layout.IsLibrary("./")

	if !ok {
		t.Errorf("Expected true, but got %t: %q", ok, err)
	}
}

func TestIsNotLibraryWhenIncludeFolderNotPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"build", "doc", "src"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	ok, err := layout.IsLibrary("./")

	if ok {
		t.Errorf("Expected false, but got %t: %q", ok, err)
	}
}

func TestIsApplicationWhenIncludeIsNotButSrcFolderPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"build", "doc", "src"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	ok, err := layout.IsApplication("./")

	if !ok {
		t.Errorf("Expected true, but got %t: %q", ok, err)
	}
}

func TestIsNotApplicationWhenBothIncludeAndSourceFoldersNotPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"build", "doc"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	ok, err := layout.IsApplication("./")

	if ok {
		t.Errorf("Expected false, but got %t: %q", ok, err)
	}
}

func TestIsNotApplicationWhenIncludeFolderPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"include", "build", "doc", "src"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	ok, err := layout.IsApplication("./")

	if ok {
		t.Errorf("Expected false, but got %t: %q", ok, err)
	}
}
