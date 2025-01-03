package controllers

import (
	"net/http"

	"github.com/Wison-cmd/GO/command/views"
)

type TemplateData struct {
	Page  string
	Title string
}

const tFile = "command/source/templates/*.html"

func RenderTemplateHome(w http.ResponseWriter, r *http.Request) {
	tName := "main"
	views.RenderTemplate(tFile, tName, w, TemplateData{Page: "Home", Title: tName})
}

func RenderTemplateForm(w http.ResponseWriter, r *http.Request) {
	tName := "form"
	views.RenderTemplate(tFile, tName, w, TemplateData{Page: "Form", Title: tName})

}
