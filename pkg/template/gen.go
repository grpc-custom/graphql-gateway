// +build ignore

package main

import (
	"flag"
	"io/ioutil"
	"os"
	"text/template"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type Params struct {
	Variable string
	Name     string
	Template string
}

const tmplFile = "./gen.tmpl"

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	name := flag.Arg(1)
	file, err := ioutil.ReadFile(filename)
	panicIf(err)
	tmpl, err := template.ParseFiles(tmplFile)
	panicIf(err)
	f, err := os.Create(filename + ".go")
	panicIf(err)
	defer f.Close()
	err = tmpl.Execute(f, &Params{
		Variable: name + "Template",
		Name:     name,
		Template: string(file),
	})
	panicIf(err)
}
