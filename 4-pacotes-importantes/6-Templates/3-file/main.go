package main

import (
	"net/http"
	"text/template"
)

type Curse struct {
	Name string
	Time int
}

type Curses []Curse

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
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
