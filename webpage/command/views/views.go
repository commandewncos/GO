package views

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Wison-cmd/GO/command/funcs"
)

type FuncMapping = map[string]any

var Builtins = FuncMapping{
	"GetYear": funcs.Functions.GetYear,
	"GetDate": funcs.Functions.GetDate,
}

func RenderTemplate(tFile, tName string, w http.ResponseWriter, data any) {
	if err := template.Must(template.New(tName).Funcs(Builtins).ParseGlob(tFile)).Execute(w, data); err != nil {
		log.Println(err.Error())
	}

}
