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

// Serve the Favicon
// TODO find a way to not have a handler function for this
func faviconHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("serving favicon")
	http.ServeFile(w, req, "public/img/favicon.ico")
}

// Future home of a web based text editor
func editorHandler(w http.ResponseWriter, req *http.Request) {
	data := Page{"GopherCode"}
	renderTemplate(w, "editor", &data)
}

// Handler for Landing/Home Page
func landingHandler(w http.ResponseWriter, req *http.Request) {
	data := Page{"GopherCode"}
	renderTemplate(w, "home", &data)
}

// Gather Templates
// TODO: progamatically gather templates
var templates = template.Must(template.ParseFiles("tmpl/editor.html", "tmpl/home.html", "tmpl/header.html", "tmpl/footer.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	templates.ExecuteTemplate(w, tmpl+".html", p)	
}

func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/editor/", editorHandler)
	http.HandleFunc("/", landingHandler)

	fmt.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
