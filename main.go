package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")

	http.HandleFunc("/home", homepageFunction)
	http.HandleFunc("/hello", helloFunction)
	http.ListenAndServe(":8080", nil)
}

func helloFunction(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello, World")
	tpl.ExecuteTemplate(w, "hello.html", nil)
}

func homepageFunction(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "homepage, r.URL.Path: %s!", r.URL.Path)
	tpl.ExecuteTemplate(w, "homepage.html", nil)
}
