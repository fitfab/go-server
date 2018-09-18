package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type page struct {
	Title  string
	Header string
	Copy   string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is Go Server!</h1>")
	fmt.Fprintf(w, "<p>go is simple as well</P>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	p := page{Title: "about", Header: "What About it", Copy: "This the copy for the about page."}
	t, _ := template.ParseFiles("basictemplate.html")
	err := t.Execute(w, p)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)
}
