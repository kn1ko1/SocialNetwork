package ui

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"socialnetwork/Server/utils"
)

var (
	Template *template.Template
)

func InitTemplates() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get current working directory:", err)
	}
	log.Println("[ui/constants] working directory is:", wd)

	// Use filepath.Join to create an absolute path
	templatesDir := filepath.Join(wd, "templates")

	// Parse the templates using ParseGlob with the absolute path
	tmpl, err := template.ParseGlob(filepath.Join(templatesDir, "*.go.html"))
	if err != nil {
		utils.HandleError("Error with template.ParseGlob in ui/init()", err)
	}
	Template = tmpl
}
