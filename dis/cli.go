package dis

import (
	"flag"
	"fmt"
	"os"

	"github.com/brettlangdon/gython/marshal"
)

func RunCommand() int {
	if len(os.Args) < 3 {
		fmt.Println("gython -dis <filename>")
		flag.PrintDefaults()
		return 1
	}

	loader, err := marshal.Load(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return 1
	}

	Disassemble(loader.Code)

	return 0
}
