package ui

import (
	"html/template"
	"log"
)

var (
	Template *template.Template
)

func init() {
	tmpl, err := template.ParseGlob("./templates/*.go.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	Template = tmpl
}
