package core

import "github.com/evilsocket/islazy/tui"

const (
	Name    = "gb"
	Version = "0.0.1"
	Author  = "Andrzej 'angelo' Lichnerowicz"
	Website = "https://andrzej.lichnerowicz.pl/"
	Desc    = "gb: Yet another build generator for cross-platform C++"
)

var (
	Banner = ""
)

func init() {
	Banner = "gb: v" + tui.Dim(Version)
}
