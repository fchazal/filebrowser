package main

import (
	"runtime"

	"github.com/fchazal/filebrowser/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
