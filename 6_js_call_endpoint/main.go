package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

/*
The point of this program is to get an input from html, have a keyup event
that reads it on javascript, get the value, send to Go, reverse the string value,
get it back on JS, and send to html
*/

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var templateFiles embed.FS

var templ = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

func main() {
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))
	http.HandleFunc("/reverse", reverse)
	http.HandleFunc("/input", input)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8989", nil)
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

// This gets a post from JS and returns the reversed string
func reverse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body := r.Body
		defer body.Close()
		// Read all the bytes from the body
		b, err := ioutil.ReadAll(body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// Parse response body
		response := struct {
			Value string
		}{}
		err = json.Unmarshal(b, &response)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		str := string(response.Value)
		// Return reverse string to w
		fmt.Fprintf(w, reverseString(str))
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
