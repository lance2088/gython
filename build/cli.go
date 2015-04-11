package build

import (
	"flag"
	"fmt"
	"os"
)

func RunCommand() int {
	if len(os.Args) < 3 {
		fmt.Println("gython -build <filename>")
		flag.PrintDefaults()
		return 1
	}

	err := Build(os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
