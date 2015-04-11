package eval

import (
	"flag"
	"fmt"
	"os"

	"github.com/brettlangdon/gython/marshal"
)

func RunCommand() int {
	if len(os.Args) < 2 {
		fmt.Println("gython <filename>")
		flag.PrintDefaults()
		return 1
	}

	loader, err := marshal.Load(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return 1
	}

	err = Eval(loader.Code)
	if err != nil {
		fmt.Println(err)
	}

	return 0
}
