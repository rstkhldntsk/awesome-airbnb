package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rstkhldntsk97/awesome-airbnb/pkg/config"
	"github.com/rstkhldntsk97/awesome-airbnb/pkg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	return mux
}
