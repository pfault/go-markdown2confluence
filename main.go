package main

import (
	"github.com/justmiles/go-markdown2confluence/cmd"
)

// version of markdown2confluence. Overwritten during build
var version = "3.4.6z"

func main() {
	cmd.Execute(version)
}
