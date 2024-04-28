package main

import (
	"os"
	"text/template"
)

type Curse struct {
	Name string
	Time int
}

type Curses []Curse

func main() {
	curse := Curse{"Go", 40}
	t := template.Must(template.New("CurseTemplate").Parse("O nome do curso é {{.Name}} e o tempo é {{.Time}}"))
	err := t.Execute(os.Stdout, curse)

	if err != nil {
		panic(err)
	}
}
