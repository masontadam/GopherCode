package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	// "io/ioutil"
	// "regexp"
)

// data to pass to templates
// TODO: port this to JSON/settings file
type Page struct {
	Title string
}

// Future home of a web based text editor
func editorHandler(w http.ResponseWriter, req *http.Request) {
	data := Page{"Text Editor"}
	renderTemplate(w, "editor", &data)
}

// Handler for Landing/Home Page
func landingHandler(w http.ResponseWriter, req *http.Request) {
	data := Page{"GopherCode"}
	renderTemplate(w, "home", &data)
}

// Handler for Contact me Page
func contactmeHandler(w http.ResponseWriter, req *http.Request) {
	data := Page{"Contact Me"}
	renderTemplate(w, "contactme", &data)
}

// Handler for Info Page
func infoHandler(w http.ResponseWriter, req *http.Request) {
	data := Page{"Info"}
	renderTemplate(w, "info", &data)
}

// Gather Templates
// TODO: progamatically gather templates
var templates = template.Must(template.ParseFiles("tmpl/editor.html",
	"tmpl/home.html", "tmpl/header.html", "tmpl/footer.html",
	"tmpl/contactme.html", "tmpl/info.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	templates.ExecuteTemplate(w, tmpl+".html", p)	
}

func main() {
	fs := http.FileServer(http.Dir("public"))

  	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.HandleFunc("/editor/", editorHandler)
	http.HandleFunc("/contactme/", contactmeHandler)
	http.HandleFunc("/info/", infoHandler)	
	http.HandleFunc("/", landingHandler)

	fmt.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
