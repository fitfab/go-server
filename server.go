package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type page struct {
	Title  string
	Header string
	Copy   string
	Aside  string
}

var tpls *template.Template

func init() {
	// use Must to handle error
	// use ParseGlob to parse all templates
	// cache all templates on this bucket "tpls"
	tpls = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func customRender(w http.ResponseWriter, templateName string, data interface{}) {
	log.Printf("templateName %v", templateName)
	err := tpls.ExecuteTemplate(w, templateName, data)
	if err != nil {
		fmt.Println(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("indexHandler %v", r.URL)
	d := page{Title: "home", Header: "Welcome home", Copy: "This is the template for the home page.", Aside: "This is extra info"}
	customRender(w, "home.tmpl", d)
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("aboutHandler %v", r.URL)
	d := page{Title: "About", Header: "What about it?", Copy: "This is the about page copy.", Aside: "This is extra for about"}
	customRender(w, "about.tmpl", d)
}

func moreHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("moreHandler %v", r.URL)
	d := page{Title: "More", Header: "This is more", Copy: "This is the more page copy.", Aside: "The data is for More with the home.tmpl"}
	customRender(w, "home.tmpl", d)
}

func main() {
	port := ":8000"
	mux := http.NewServeMux()

	// setup to serve static assets: css, js and images
	// good video: https://www.youtube.com/watch?v=fz8pcJTLntI explaining "http.StripPrefix"
	fs := http.StripPrefix("/static", http.FileServer(http.Dir("./static")))
	mux.Handle("/static/", fs)

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/more", moreHandler)

	log.Printf("\nListenting at %v", port)
	http.ListenAndServe(":8000", mux)
}
