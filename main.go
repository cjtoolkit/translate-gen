package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/cjtoolkit/translate-gen/structure"
	"github.com/cjtoolkit/translate-gen/template"
)

func main() {
	if len(os.Args) <= 2 {
		return
	}

	packageName := os.Args[1]

	bases := []template.Context{}

	for _, fileName := range os.Args[2:] {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		fileBase := structure.FileBase{}
		_, err = toml.DecodeReader(file, &fileBase)
		if err != nil {
			panic(err)
		}

		bases = append(bases, template.Context{
			Package:  packageName,
			Source:   fileName,
			FileBase: fileBase,
		})
		file.Close()
	}

	aTemplate := template.BuildTemplate()
	for _, base := range bases {
		splitName := strings.Split(base.Source, ".")
		name := strings.Join(splitName[:len(splitName)-1], ".") + ".go"

		file, err := os.Create(name)
		if err != nil {
			panic(err)
		}

		err = aTemplate.Execute(file, base)
		if err != nil {
			panic(err)
		}

		exec.Command("go", "fmt", name).Run()
	}
}
