package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

type page struct {
	Title  string
	Header string
	Copy   string
}

var opts = renderer.Options{
	ParseGlobPattern: "./templates/*.html",
}
var rnd = renderer.New(opts)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	d := page{Title: "home", Header: "Welcome home", Copy: "This is run by."}
	err := rnd.HTML(w, http.StatusOK, "layout", d)
	if err != nil {
		fmt.Println(err)
	}
}

// using template from golang
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	p := page{Title: "about", Header: "What About it", Copy: "This the copy for the about page."}
	t, _ := template.ParseFiles("./templates/index.html")
	err := t.Execute(w, p)
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
