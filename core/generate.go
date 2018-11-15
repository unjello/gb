package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bmatcuk/doublestar"
	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/islazy/tui"
)

func checkIfSourceFolderExists(project_root string) error {
	path := filepath.Join(project_root, "src")
	fi, err := os.Stat(path)
	if err != nil {
		log.Warning("Source folder " + tui.Green("src") + " not found in project root. You should create one.")
		return err
	}
	if mode := fi.Mode(); mode.IsDir() != true {
		log.Warning(tui.Green("src") + " found in project root, but it is not a folder.")
		return fmt.Errorf("Source folder is not a directory")
	}

	return nil
}

func checkIfIncludeFolderExists(project_root string) error {
	path := filepath.Join(project_root, "include")
	fi, err := os.Stat(path)
	if err != nil {
		log.Warning("Include folder " + tui.Green("include") + " not found in project root. You should create one.")
		return err
	}
	if mode := fi.Mode(); mode.IsDir() != true {
		log.Warning(tui.Green("include") + " found in project root, but it is not a folder.")
		return fmt.Errorf("Include folder is not a directory")
	}

	return nil
}

func checkIfBuildFolderIsIgnored(project_root string) error {
	path := filepath.Join(project_root, ".gitignore")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Warning("No " + tui.Green(".gitignore") + " in project root. You should create one.")
		return err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Could not read .gitignore contents")
	}
	r, _ := regexp.Compile("(?m)^build/?$")
	match := r.Match(b)
	if match == false {
		log.Warning("Your build folder " + tui.Green("build/") + " should be ignored by git. Add it to " + tui.Green(".gitignore"))
	}
	return nil
}

func ensureBuildFolderExists(project_root string) (string, error) {
	path := filepath.Join(project_root, "build")
	log.Debug("Creating build folder " + tui.Dim(path))
	err := os.MkdirAll(path, os.ModeDir|0775)
	if err != nil {
		log.Error("Failed to create build folder " + tui.Dim(path))
		return "", err
	}

	return path, nil
}

func generateConanDependencies(build_root string) error {
	conanFileTemplate := `
[requires]
doctest/2.0.0@unjello/testing

[generators]
json
`
	conanFile, err := ExecuteTemplate("conanFileTemplate", conanFileTemplate, nil)
	if err != nil {
		return err
	}

	path := filepath.Join(build_root, "conanfile.txt")
	log.Debug("Generating conan dependency file " + tui.Dim(path))
	err = ioutil.WriteFile(path, []byte(conanFile), 0664)
	if err != nil {
		log.Error("Failed to create a file " + tui.Green(path) + "\n" + tui.Red(err.Error()))
		return err
	}

	runner := OsCommandRunner{}
	runner.RunWithOutput([]string{"conan", "install", "-if", "build/", "build/"})

	return nil
}

type SourceFile struct {
	FullPath  string
	RelPath   string
	BaseName  string
	Extension string
}

func getSourceFiles(sourceRoot string, buildRoot string) ([]SourceFile, error) {
	sourceGlob := filepath.Join(sourceRoot, "**/*.cpp")
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

func generateNinjaBuildFile(projectRoot string, buildRoot string, projectName string, isHeaderOnly bool) error {
	type BuildInfo struct {
		Name           string
		PublicIncludes string
		Sources        []SourceFile
		Tests          []SourceFile
		TestsIncludes  []string
	}
	ninjaBuildTemplate := `
ninja_required_version = 1.3

cxx = g++-8
builddir = out
testsbuilddir = out/tests
bindir = bin/
testbindir = $bindir/tests

cxxflags = -Wall -Werror -std=c++17
ldflags = -L$builddir
testcxxflags = {{range .TestsIncludes}}-I {{.}} {{end}}


rule cxx
  command = $cxx $cxxflags -c ${in} -o ${out}
  description = CXX $out
  depfile = $out.d
  deps = gcc

rule testcxx
  command = $cxx $cxxflags $testcxxflags -c ${in} -o ${out}
  description = CXX $out
  depfile = $out.d
  deps = gcc

rule link
  command = $cxx $linkflags $in -o $out
  description = LINK $out

{{range .Sources}}
build $builddir/{{.BaseName}}.o: cxx {{.RelPath}}
{{end}}

{{range .Tests}}
build $testsbuilddir/{{.BaseName}}.o: testcxx {{.RelPath}}
{{end}}

build $bindir/{{.Name}}: link {{range .Sources}}$builddir/{{.BaseName}}.o {{end}}
{{range .Tests}}
build $testbindir/{{.BaseName}}: link $testsbuilddir/{{.BaseName}}.o
{{end}}

build all: phony $bindir/{{.Name}} {{range .Tests}}$testbindir/{{.BaseName}} {{end}}

`
	ninjaTestsOnlyBuildTemplate := `
ninja_required_version = 1.3

cxx = g++-8
builddir = out
testsbuilddir = out/tests
bindir = bin/
testbindir = $bindir/tests

cxxflags = -Wall -Werror -std=c++17
ldflags = -L$builddir
testcxxflags = {{range .TestsIncludes}}-I {{.}} {{end}} -I {{.PublicIncludes}}


rule testcxx
  command = $cxx $cxxflags $testcxxflags -c ${in} -o ${out}
  description = CXX $out
  depfile = $out.d
  deps = gcc

rule link
  command = $cxx $linkflags $in -o $out
  description = LINK $out

{{range .Tests}}
build $testsbuilddir/{{.BaseName}}.o: testcxx {{.RelPath}}
build $testbindir/{{.BaseName}}: link $testsbuilddir/{{.BaseName}}.o
{{end}}


build all: phony {{range .Tests}}$testbindir/{{.BaseName}} {{end}}

`

	info, err := ReadConanBuildInfo(filepath.Join(buildRoot, "conanbuildinfo.json"))
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	doctest, err := GetTestingPackage(info)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	sources, err := getSourceFiles(filepath.Join(projectRoot, "src"), buildRoot)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	tests, err := getSourceFiles(filepath.Join(projectRoot, "test"), buildRoot)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	buildInfo := BuildInfo{
		projectName,
		filepath.Join(projectRoot, "include"),
		sources,
		tests,
		doctest.IncludePaths,
	}

	var ninjaFile string
	if isHeaderOnly == true {
		var err error
		ninjaFile, err = ExecuteTemplate("ninjaBuildTemplate", ninjaTestsOnlyBuildTemplate, buildInfo)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
	} else {
		var err error
		ninjaFile, err = ExecuteTemplate("ninjaBuildTemplate", ninjaBuildTemplate, buildInfo)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
	}

	path := filepath.Join(buildRoot, "build.ninja")
	log.Debug("Generating ninja build file " + tui.Dim(path))
	err = ioutil.WriteFile(path, []byte(ninjaFile), 0664)
	if err != nil {
		log.Error("Failed to create a file " + tui.Green(path) + "\n" + tui.Red(err.Error()))
		return err
	}
	return nil
}

func GenerateBuildScripts() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory")
	}
	log.Info("Generating build for project dir: " + tui.Dim(projectRoot))
	checkIfBuildFolderIsIgnored(projectRoot)
	srcErr := checkIfSourceFolderExists(projectRoot)
	includeErr := checkIfIncludeFolderExists(projectRoot)
	isHeaderOnly := srcErr != nil && includeErr == nil

	buildRoot, _ := ensureBuildFolderExists(projectRoot)

	projectName := filepath.Base(projectRoot)
	log.Info("Infering project name from folder: " + tui.Green(projectName))

	generateConanDependencies(buildRoot)
	generateNinjaBuildFile(projectRoot, buildRoot, projectName, isHeaderOnly)
}
