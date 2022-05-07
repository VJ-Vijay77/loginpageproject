package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
     HttpOnly: true,
	 Path: "/",
	 Secure: false,
	 SameSite: http.SameSiteLaxMode,
    })
	return csrfHandler
}

func SessionLoad(next http.Handler)http.Handler{
	return session.LoadAndSave(next)
}