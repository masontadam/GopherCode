// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"fmt"
	"log"
	"net/http"
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
var templates = template.Must(template.ParseFiles("tmpl/editor.html", "tmpl/home.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	fmt.Println(tmpl)
	templates.ExecuteTemplate(w, tmpl+".html", p)	
}

func main() {
	http.HandleFunc("/editor/", editorHandler)
	http.HandleFunc("/", landingHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
