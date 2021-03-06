package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/evilsocket/islazy/log"
)

type Dependency struct {
	RootPath     string
	Version      string
	IncludePaths []string `json:"include_paths"`
	Name         string
	Description  string
}
type Setting struct {
	Architecture    string `json:"arch"`
	OsBuild         string `json:"os_build"`
	Os              string
	Compiler        string
	CompilerVersion string `json:"compiler.version"`
	CompilerLibCxx  string `json:"compiler.libcxx"`
}
type ConanBuildInfo struct {
	Dependencies []Dependency
	Settings     Setting
}

type ConanPackage struct {
	Name    string
	Version string
	Channel string
}

func ReadConanBuildInfo(path string) (ConanBuildInfo, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
		return ConanBuildInfo{}, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var buildInfo ConanBuildInfo

	json.Unmarshal(byteValue, &buildInfo)
	return buildInfo, nil
}

func GetTestingPackage(info ConanBuildInfo) (Dependency, error) {
	for _, v := range info.Dependencies {
		if v.Name == "doctest" {
			return v, nil
		}
	}
	return Dependency{}, fmt.Errorf("Not found")
}

func ParsePackageString(name string) (ConanPackage, error) {
	var re = regexp.MustCompile(`(?m)([^\/]+)\/(\d+\.\d+\.\d+)\@([^\/]+\/[^$]+)`)
	info := re.FindStringSubmatch(name)
	if info == nil {
		return ConanPackage{}, fmt.Errorf("Could not parse conan package string")
	}

	return ConanPackage{info[1], info[2], info[3]}, nil
}
