package main

import (
	"os"
	"text/template"
)

type Curse struct {
	Name string
	Time int
}

func main() {
	curse := Curse{"Go", 40}
	tmp := template.New("CurseTemplate")
	tmp, _ = tmp.Parse("O nome do curso é {{.Name}} e o tempo é {{.Time}}")
	err := tmp.Execute(os.Stdout, curse)

	if err != nil {
		panic(err)
	}
}
