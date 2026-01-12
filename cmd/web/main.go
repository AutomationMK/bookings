package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AutomationMK/bookings/internal/config"
	"github.com/AutomationMK/bookings/internal/handlers"
	"github.com/AutomationMK/bookings/internal/models"
	"github.com/AutomationMK/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	// specify special structs that will be stored in the session
	gob.Register(models.Reservation{})
	// initialize a session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// give render acess to our app config var
	render.NewTemplate(&app)

	fmt.Printf("Starting application on port %s\n", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
