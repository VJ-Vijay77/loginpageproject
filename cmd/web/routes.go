package main

import (
	"net/http"

	"github.com/VJ-Vijay77/LoginPageNew/pkg/config"
	"github.com/VJ-Vijay77/LoginPageNew/pkg/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig)http.Handler{
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/",handler.Repo.Home)
	mux.Get("/abt",handler.Repo.About)
	return mux
}