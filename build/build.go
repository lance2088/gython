package build

import (
	"fmt"
	"os"
	"path"

	"github.com/jteeuwen/go-bindata"
)

func Build(filename string, dir string) error {
	config := bindata.NewConfig()

	config.Output = path.Join(dir, "bindata.go")
	config.Input = []bindata.InputConfig{
		bindata.InputConfig{
			Path:      filename,
			Recursive: false,
		},
	}

	err := bindata.Translate(config)
	if err != nil {
		return err
	}

	fp, err := os.Create(path.Join(dir, "main.go"))
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fmt.Fprintf(fp, `package main

import (
    "fmt"

    "github.com/brettlangdon/gython/eval"
    "github.com/brettlangdon/gython/marshal"
)

func main() {
    data, err := Asset("%s")
    if err != nil {
        fmt.Println(err)
        return
    }

    loader, err := marshal.LoadString(data, "%s")
    if err != nil {
        fmt.Println(err)
        return
    }

    err = eval.Eval(loader.Code)
    if err != nil {
        fmt.Println(err)
    }
}`, filename, filename)
	if err != nil {
		return err
	}

	return nil
}
