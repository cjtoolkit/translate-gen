package main

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/cjtoolkit/translate-gen/structure"
	"github.com/cjtoolkit/translate-gen/template"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) <= 2 {
		return
	}

	packageName := os.Args[1]

	bases := []template.Context{}

	for _, fileName := range os.Args[2:] {
		file, err := os.Open(fileName)
		checkErr(err)

		fileBase := structure.FileBase{}
		_, err = toml.DecodeReader(file, &fileBase)
		checkErr(err)

		bases = append(bases, template.Context{
			Package:  packageName,
			Source:   fileName,
			FileBase: fileBase,
		})
		file.Close()
	}

	aTemplate := template.BuildTemplate()
	for _, base := range bases {
		splitName := strings.Split(path.Base(base.Source), ".")
		name := strings.Join(splitName[:len(splitName)-1], ".") + ".go"

		file, err := os.Create(name)
		checkErr(err)

		err = aTemplate.Execute(file, base)
		checkErr(err)
		file.Close()

		err = exec.Command("go", "fmt", name).Run()
		checkErr(err)
	}
}
