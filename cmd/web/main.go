package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VJ-Vijay77/LoginPageNew/pkg/config"
	"github.com/VJ-Vijay77/LoginPageNew/pkg/handler"
	"github.com/VJ-Vijay77/LoginPageNew/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	

	app.InProduction = false 

	session = scs.New()
	session.Lifetime = 24 *time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln(err)
	}
	app.TemplateCache = tc
	app.UseCache = false 

	render.NewTemplate(&app)

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)
	
	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/page", handler.Repo.InPage)
	fmt.Println(fmt.Sprintf("Starting port number %s", port))
	// http.ListenAndServe(port, nil)
	srv := &http.Server{
		Addr: port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatalln(err)
}


