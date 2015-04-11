package main

import (
	"fmt"
	"os"

	"github.com/brettlangdon/gython/dis"
	"github.com/brettlangdon/gython/marshal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go run main.go <filename>")
		return
	}

	loader, err := marshal.Load(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	dis.Disassemble(loader.Code)
}
