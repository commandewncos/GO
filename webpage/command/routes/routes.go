package routes

import (
	"net/http"

	"github.com/Wison-cmd/GO/command/controllers"
)

func RequestHandler() http.Handler {
	mux := http.NewServeMux()

	// Server static file: CSS, Javascript and Assets
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("command/source/static"))))

	mux.HandleFunc("/", controllers.RenderTemplateHome)
	mux.HandleFunc("/form", controllers.RenderTemplateForm)

	return mux
}
