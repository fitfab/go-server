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
	// using Must to handle error and ParseGlob
	// cache all templates on this bucket "tpls"
	tpls = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func customRender(w http.ResponseWriter, name string, data interface{}) {
	err := tpls.ExecuteTemplate(w, "layout.tmpl", data)
	if err != nil {
		fmt.Println(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	d := page{Title: "home", Header: "Welcome home", Copy: "This is run by.", Aside: "This is extra info"}
	customRender(w, "layout.tmpl", d)
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	d := page{Title: "About", Header: "What about it?", Copy: "This is the about page copy."}
	customRender(w, "layout.tmpl", d)
}

func moreHandler(w http.ResponseWriter, r *http.Request) {
	d := page{Title: "More", Header: "This is more", Copy: "This is the more page copy."}
	customRender(w, "layout.tmpl", d)
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
	log.Printf("\nlistenting at %v", port)
	http.ListenAndServe(":8000", mux)
}
