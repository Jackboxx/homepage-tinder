package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &WebsitePanel{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "homepage tinder",
		Description: "Tinder swiping through a list of websites homepage's",
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
