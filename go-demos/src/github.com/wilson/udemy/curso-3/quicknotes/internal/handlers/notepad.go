package handlers

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Wilson-cmd/quicknotes/internal/handlers/custom"
)

type notepadHandler struct {
}

func NewNotepadHandler() *notepadHandler {
	return &notepadHandler{}
}

// Declare var files: path files
var (
	files = struct {
		base, header, home, notepad, view, footer string
	}{
		base:    "views/templates/pages/base.html",
		header:  "views/templates/pages/header.html",
		home:    "views/templates/pages/home.html",
		notepad: "views/templates/pages/notepad.html",
		view:    "views/templates/pages/view.html",
		footer:  "views/templates/pages/footer.html",
	}
)

func (nh *notepadHandler) Home(w http.ResponseWriter, r *http.Request) {

	//Coringa
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	t, e := template.ParseFiles(files.base, files.header, files.home, files.footer)

	if e != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	t.ExecuteTemplate(w, "base", nil)

}

func (nh *notepadHandler) Notepad(w http.ResponseWriter, r *http.Request) error {

	t, e := template.ParseFiles(files.base, files.header, files.notepad, files.footer)
	if e != nil {
		return errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	return t.ExecuteTemplate(w, "base", nil)

}

func (nh *notepadHandler) NotepadView(w http.ResponseWriter, r *http.Request) error {

	id := r.URL.Query().Get("id")

	if id == "" {
		return custom.WithStatusError(
			errors.New("required value"),
			http.StatusBadRequest,
		)
	}

	t, e := template.ParseFiles(files.base, files.header, files.view, files.footer)
	if e != nil {
		return custom.WithStatusError(
			errors.New("internal error"),
			http.StatusInternalServerError,
		)

	}

	return t.ExecuteTemplate(w, "base", id)

}

func (nh *notepadHandler) NotepadCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// Operation deny
		http.Error(w, "Ops!...", http.StatusMethodNotAllowed)
		return

	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Notecreate..."))

	}
}
