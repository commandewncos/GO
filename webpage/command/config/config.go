package config

import (
	"flag"
	f "fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/Wison-cmd/GO/command/routes"
)

type Config struct{ Port, Env, Version string }

func (c *Config) CommandsFlag() {
	c.Version = "0.0.1"
	flag.StringVar(&c.Port, "port", "3000", "define server port")
	flag.StringVar(&c.Env, "environment", "dev", "define environment")
	flag.Parse()
}

type Application struct {
	//Name Type
	Config Config
	Cache  map[string]*template.Template
}

func (app *Application) Start() error {
	server := &http.Server{

		Addr:              f.Sprintf(":%s", app.Config.Port),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           routes.RequestHandler(),
	}
	return server.ListenAndServe()
}
