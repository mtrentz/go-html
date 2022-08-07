package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Pessoa struct {
	Nome      string
	Idade     int
	Profissao string
}

var templ = template.Must(template.ParseFiles("template.html"))

func main() {
	http.HandleFunc("/", index)
	fmt.Println("serving on 8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p := Pessoa{
		Nome:      "João",
		Idade:     30,
		Profissao: "Programador",
	}
	// templ.Execute(w, p)
	templ.ExecuteTemplate(w, "Index", p)
}
