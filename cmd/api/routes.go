package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// sets routes as method on Config struct
// so if Config has DB string and we create app as instance of Config with DB set equal to hello world
// And routes was a method that prints app.DB
// app.routes() would print hello world
// attached to a type using receiver, thus a method
func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// check server is alive
	mux.Use(middleware.Heartbeat("/ping"))

	// on post request to /, calls broker method
	mux.Post("/", app.Broker)

	return mux
}
