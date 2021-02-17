package vesper

import (
	"flag"
	"fmt"
	"os"
)

// Extension allows for extending Vesper
type Extension interface {
	Init() error
	Cleanup()
	String() string
}

var extensions []Extension

// Version - this version of vesper
const Version = "vesper v0.1"

// Main entrypoint for main Vesper interpreter
func Main() {
	var help, version bool
	var path string
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&version, "version", false, "Show the current version")
	flag.StringVar(&path, "path", "", "add directories to vesper load path")

	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(1)
	}

	if version {
		fmt.Println(Version)
		os.Exit(0)
	}
	args := flag.Args()

	if len(args) > 0 {
		Run(args...)
	} else {
		REPL()
	}

	Cleanup()
}

// Run executes the requested files
func Run(args ...string) {
}

// REPL starts a REPL
func REPL() {
}

// Cleanup deinitialise the runtime
func Cleanup() {
}
