package main

import (
	"net/http"
	"strings"
	"text/template"
)

type Curse struct {
	Name string
	Time int
}

type Curses []Curse

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))
		// t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, Curses{
			{"Go", 40},
			{"Python", 20},
			{"Java", 10},
			{"JavaScript", 5},
		})

		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":3000", nil)

}
