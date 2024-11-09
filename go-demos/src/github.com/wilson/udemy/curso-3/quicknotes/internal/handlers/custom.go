package handlers

import (
	"errors"
	"net/http"

	"github.com/Wilson-cmd/quicknotes/internal/handlers/custom"
)

type HandleWithError func(w http.ResponseWriter, r *http.Request) error

func (f HandleWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := f(w, r)
	if e != nil {
		var status custom.StatusError
		if errors.As(e, &status) {
			http.Error(w, e.Error(), status.StatusCustomHttpError())
			return
		}

		http.Error(w, e.Error(), http.StatusInternalServerError)
	}

}
