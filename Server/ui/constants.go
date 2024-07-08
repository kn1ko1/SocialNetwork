package ui

import (
	"html/template"
	"socialnetwork/Server/utils"
)

var (
	Template *template.Template
)

func InitTemplates() {

	// Parse the templates using ParseGlob
	tmpl, err := template.ParseGlob("../Server/templates/*.go.html")
	if err != nil {
		utils.HandleError("Error with template.ParseGlob in ui/init()", err)
	}
	Template = tmpl
}
