package main

import (
	"net/http"
	"text/template"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

type Aluno struct {
	Nome  string
	Idade int
}

type Escola struct {
	Nome      string
	Alunos    []Aluno
	NumAlunos int
}

var alunos []Aluno
var escola Escola

func main() {
	alunos = []Aluno{
		{"João", 10},
		{"Maria", 12},
		{"Pedro", 15},
	}

	escola = Escola{
		Nome:      "Escola 1",
		Alunos:    alunos,
		NumAlunos: len(alunos),
	}

	http.HandleFunc("/", school)
	http.HandleFunc("/students", students)
	http.ListenAndServe(":8080", nil)

}

func school(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "School", escola)
}

func students(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "Students", alunos)
}
