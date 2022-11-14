package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/romimusic/bookingSystem/pkg/config"
	"github.com/romimusic/bookingSystem/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

//this is a external package to handle the templates the repo is https://github.com/bmizerany/pat
// mux := pat.New()

// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
