package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var templ = template.Must(template.ParseGlob("templates/*.gohtml"))

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
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
		// fmt.Println(r)
		fmt.Println(r.FormValue("nome"))
		// Redirect home
		http.Redirect(w, r, "/", 301)
	}
}
