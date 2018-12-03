package core

import (
	"github.com/unjello/gb/layout"
)

type BuildInfo struct {
	Cxx        string
	CxxOptions string
	Project    layout.ProjectInfo
}

func NewBuildInfo() BuildInfo {
	return BuildInfo{}
}
