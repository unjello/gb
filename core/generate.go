package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/islazy/tui"
	"github.com/unjello/gb/layout"
)

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

func generateConanDependencies(build_root string, dependencies []string) error {
	conanFileTemplate := `
[requires]{{range .}}
{{ . }}
{{end}}
[generators]
json
`
	conanFile, err := ExecuteTemplate("conanFileTemplate", conanFileTemplate, dependencies)
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

func generateNinjaBuildFile(projectRoot string, buildRoot string, projectName string, project layout.ProjectInfo) error {
	ninjaBuildTemplate := `
ninja_required_version = 1.3

cxx = g++-8
builddir = out
testsbuilddir = out/tests
bindir = bin/
testbindir = $bindir/tests

cxxflags = -Wall -Werror -std=c++2a -fconcepts
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

cxxflags = -Wall -Werror -std=c++2a -fconcepts
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
build $testsbuilddir/{{getTestName .}}.o: testcxx {{.RelPath}}
build $testbindir/{{getTestName .}}: link $testsbuilddir/{{getTestName .}}.o
{{end}}


build all: phony {{range .Tests}}$testbindir/{{getTestName .}} {{end}}

`
	buildInfo := project
	buildInfo.Name = projectName

	if buildInfo.Dependencies != nil {
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

		buildInfo.TestsIncludes = doctest.IncludePaths
	}

	var ninjaFile string
	if project.Type == layout.HeaderOnly {
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
	err := ioutil.WriteFile(path, []byte(ninjaFile), 0664)
	if err != nil {
		log.Error("Failed to create a file " + tui.Green(path) + "\n" + tui.Red(err.Error()))
		return err
	}
	return nil
}

func GenerateBuildScripts() {
	if err := VerifyConanExists(); err != nil {
		log.Fatal("Could not find Conan. Please install it from https://conan.io")
	}

	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory")
	}
	log.Info("Generating build for project dir: " + tui.Dim(projectRoot))
	checkIfBuildFolderIsIgnored(projectRoot)
	buildRoot, _ := ensureBuildFolderExists(projectRoot)

	projectLayout := layout.NewDefaultProjectLayout()
	project, err := projectLayout.Get(projectRoot, buildRoot)
	if err != nil {
		log.Fatal("Failed to understand project structure")
	}

	projectName := filepath.Base(projectRoot)
	log.Info("Infering project name from folder: " + tui.Green(projectName))

	if project.Dependencies != nil {
		fmt.Println(project.Dependencies)
		generateConanDependencies(buildRoot, project.Dependencies)
	}
	generateNinjaBuildFile(projectRoot, buildRoot, projectName, project)
}
