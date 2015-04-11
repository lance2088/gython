package main

import (
	"flag"
	"os"

	"github.com/brettlangdon/gython/build"
	"github.com/brettlangdon/gython/dis"
	"github.com/brettlangdon/gython/eval"
)

var runDis bool
var runBuild bool

func main() {
	flag.BoolVar(&runDis, "dis", false, "disassemble the file and exit")
	flag.BoolVar(&runBuild, "build", false, "build the file and exit")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var res = 0
	if runDis {
		res = dis.RunCommand()
	} else if runBuild {
		res = build.RunCommand()
	} else {
		eval.RunCommand()
	}

	os.Exit(res)
}
