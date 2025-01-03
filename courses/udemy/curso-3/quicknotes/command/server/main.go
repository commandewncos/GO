package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Wilson-cmd/quicknotes/internal/handlers"
	"github.com/Wilson-cmd/quicknotes/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	settings := loadSettings()
	logger := newLogger(os.Stderr, settings.GetLevelLog())

	//Set default logger like logger from function newLogger
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	pgxConnPool, err := pgxpool.New(context.Background(), settings.DatabaseConnection)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("Server running on port %s\n", settings.ServerPort))
	defer pgxConnPool.Close()

	staticHandler := http.FileServer(http.Dir("views/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	notepadHandler := handlers.NewNotepadHandler()

	mux.HandleFunc("/", notepadHandler.Home)

	noteRepo := repositories.NewNoteRepository(pgxConnPool)

	var (
		id                    int
		title, content, color string
	)
	id = 2
	title = "Note UPDATED"
	content = "This content not empty"
	color = "#45FAFF"

	note, err := noteRepo.UpdateNote(id, title, content, color)
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println(note)

	mux.Handle("/notepad", handlers.HandleWithError(notepadHandler.Notepad))
	mux.Handle("/notepad/view", handlers.HandleWithError(notepadHandler.NotepadView))

	mux.HandleFunc("/notepad/create", notepadHandler.NotepadCreate)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", settings.ServerPort), mux); err != nil {
		panic(err)
	}

}
