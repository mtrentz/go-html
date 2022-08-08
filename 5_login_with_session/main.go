package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/didip/tollbooth/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var templateFiles embed.FS

var templ = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

// Single user/password, defined here
var adminUser = "admin"
var adminPass = "admin"

// Session
var store = sessions.NewCookieStore([]byte("secret-should-go-here"))

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(staticFiles)))
	// 1 request per second on all endpoints
	r.Handle("/home", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil), indexHandler))
	r.Handle("/login", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil), loginHandler))
	r.Handle("/loginauth", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil), loginAuthHandler))
	err := http.ListenAndServe(":8989", r)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	_, ok := session.Values["authenticated"]
	if ok {
		templ.ExecuteTemplate(w, "Index", nil)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "Login", nil)
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == adminUser && password == adminPass {
		session, _ := store.Get(r, "user-session")
		session.Options = &sessions.Options{
			Path: "/",
			// 1 hour
			MaxAge: 3600,
		}
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
