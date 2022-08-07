package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var templateFiles embed.FS

var templ = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

func main() {
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))
	http.HandleFunc("/", index)
	http.HandleFunc("/input", input)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "Index", nil)
}

// Takes in post that has action to "input" and prints to stdout
func input(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println(r.FormValue("nome"))
		http.Redirect(w, r, "/", 301)
	}
}
