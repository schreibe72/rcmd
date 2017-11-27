package main

import "github.com/schreibe72/rcmd/cmd"

var (
	version string
	githash string
)

func main() {
	cmd.Version = version
	cmd.Githash = githash
	cmd.Execute()
}
