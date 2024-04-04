package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/a-h/templdockertest/templates"
)

func main() {
	http.Handle("/", templ.Handler(templates.Hello()))
	http.ListenAndServe(":8080", nil)
}
