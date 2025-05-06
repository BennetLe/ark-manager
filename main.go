package main

import (
	"ark-manager/view"
	"ark-manager/view/layout"
	"ark-manager/view/partial"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	c := layout.Base(view.Index())
	http.Handle("/", templ.Handler(c))

	http.Handle("/foo", templ.Handler(partial.Foo()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
