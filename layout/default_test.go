package layout

import (
	"testing"

	"github.com/evilsocket/islazy/log"
	"github.com/spf13/afero"
)

func init() {
	log.Level = log.FATAL + 1
}

func TestIsHeaderOnlyLibraryWhenOnlyIncludeFolderPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"include", "build", "doc"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	meta, err := layout.Get("./", "./build")

	if err != nil {
		t.Errorf("Expected no error, but got: %q", err)
	}

	if meta.Type != HeaderOnly {
		t.Errorf("Expected project to be HeaderOnly, but got: %q", meta.Type)
	}
}

func TestIsLibraryWhenIncludeAndSrcFoldersPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"include", "build", "doc", "src"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	meta, err := layout.Get("./", "./build")

	if err != nil {
		t.Errorf("Expected no error, but got: %q", err)
	}

	if meta.Type != Library {
		t.Errorf("Expected project to be Library, but got: %q", meta.Type)
	}
}

func TestIsApplicationWhenIncludeIsNotButSrcFolderPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"build", "doc", "src"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	meta, err := layout.Get("./", "./build")

	if err != nil {
		t.Errorf("Expected no error, but got: %q", err)
	}

	if meta.Type != Application {
		t.Errorf("Expected project to be Application, but got: %q", meta.Type)
	}
}

func TestIsNotApplicationWhenBothIncludeAndSourceFoldersNotPresent(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	for _, folder := range []string{"build", "doc"} {
		AppFs.MkdirAll(folder, 0755)
	}

	layout := NewDefaultProjectLayout()
	_, err := layout.Get("./", "./build")

	if err == nil {
		t.Errorf("Expected an error, but got: %q", err)
	}
}
