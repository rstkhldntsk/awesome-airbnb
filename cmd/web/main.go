package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/rstkhldntsk97/awesome-airbnb/pkg/config"
	"github.com/rstkhldntsk97/awesome-airbnb/pkg/handler"
	"github.com/rstkhldntsk97/awesome-airbnb/pkg/render"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTmplCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)
	fmt.Println("Starting application on port ", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
