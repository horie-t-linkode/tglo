package main

import (
	"tglo_cli/subcommand"
)

var version = "unknown"
var revision = "unknown"
var buildDate = "unknown"

func main() {
	subcommand.Execute(version, revision, buildDate)
}