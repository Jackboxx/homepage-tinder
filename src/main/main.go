package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &WebsitePanel{})
	app.RunWhenOnBrowser()

	h := app.Handler{
		Author:      "Moritz Gschwantner",
		Name:        "Business Tinder",
		Description: "A Tinder like interface to view a list of companies",
		Keywords: []string{
			"tinder",
			"websites",
			"business",
		},
		Styles: []string{
			"web/main.css",
			"https://fonts.googleapis.com/css2?family=Poppins&display=swap",
		},
	}

	http.Handle("/", &app.Handler{
		Name:        "homepage tinder",
		Description: "Tinder swiping through a list of websites homepage's",
	})

	if err := http.ListenAndServe(":5050", &h); err != nil {
		log.Fatal(err)
	}
}
