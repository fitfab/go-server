package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

type page struct {
	Title  string
	Header string
	Copy   string
	Aside  string
}

var rnd *renderer.Render

func init() {
	rnd = renderer.New()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	d := page{Title: "home", Header: "Welcome home", Copy: "This is run by."}
	tpls := []string{"templates/layout.tmpl", "templates/home.tmpl", "templates/partial.tmpl"}
	err := rnd.Template(w, http.StatusOK, tpls, d)
	if err != nil {
		fmt.Println(err)
	}
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {

	d := page{Title: "About", Header: "What about this?", Copy: "This is run by renderer.", Aside: "This is an example using Template method. "}
	tpls := []string{"templates/layout.tmpl", "templates/about.tmpl", "templates/partial.tmpl"}
	err := rnd.Template(w, http.StatusOK, tpls, d)
	if err != nil {
		fmt.Println(err)
	}
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
	log.Printf("\nlistenting at %v", port)
	http.ListenAndServe(":8000", mux)
}
