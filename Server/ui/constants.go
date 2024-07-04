package ui

import (
	"html/template"
	"path/filepath"
	"socialnetwork/Server/utils"
)

var (
	Template *template.Template
)

func InitTemplates() {
	// Get the absolute path to the templates directory
	templatesDir := filepath.Join("..", "templates", "*.go.html")

	// Parse the templates using ParseGlob
	tmpl, err := template.ParseGlob(templatesDir)
	if err != nil {
		utils.HandleError("Error with template.ParseGlob in ui/init()", err)
	}
	Template = tmpl
}
