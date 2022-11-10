package main

import (
	"runtime"

	cmd "github.com/anonymitycash/anonymitycash/cmd/anonymitycashcli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
