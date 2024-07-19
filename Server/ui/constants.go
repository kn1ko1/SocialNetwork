package ui

import (
	"html/template"
	dbUtils "socialnetwork/Database/databaseUtils"
	"socialnetwork/Server/utils"
)

var (
	Template *template.Template
)

func InitTemplates() {

	var templateString string

	if dbUtils.RunningInDocker() {
		templateString = "Server/templates/*.go.html"
	} else {
		templateString = "../Server/templates/*.go.html"
	}

	// Parse the templates using ParseGlob
	tmpl, err := template.ParseGlob(templateString)
	if err != nil {
		utils.HandleError("Error with template.ParseGlob in ui/init()", err)
	}
	Template = tmpl
}
