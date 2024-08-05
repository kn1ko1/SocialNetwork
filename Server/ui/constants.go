package ui

import (
	"html/template"
<<<<<<< HEAD
	dbUtils "socialnetwork/Database/databaseUtils"
=======
	"log"
	"os"
	"path/filepath"
>>>>>>> main
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

<<<<<<< HEAD
	var templateString string

	if dbUtils.RunningInDocker() {
		templateString = "Server/templates/*.go.html"
	} else {
		templateString = "../Server/templates/*.go.html"
	}

	// Parse the templates using ParseGlob
	tmpl, err := template.ParseGlob(templateString)
=======
	// Parse the templates using ParseGlob with the absolute path
	tmpl, err := template.ParseGlob(filepath.Join(templatesDir, "*.go.html"))
>>>>>>> main
	if err != nil {
		utils.HandleError("Error with template.ParseGlob in ui/init()", err)
	}
	Template = tmpl
}
