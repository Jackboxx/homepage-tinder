package main

import (
	"log"
	"math/rand"
	"net/http"
	"reflect"

	"homepage-tinder/src/resources"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type website struct {
	app.Compo
	Name     string
	ImageSrc string
}

func (ws *website) Render() app.UI {
	element := readRandomElement()
	return app.Div().Text("some text: " + element.Url)
}

func readRandomElement() resources.Data {
	keys := reflect.ValueOf(resources.WebsiteData).MapKeys()
	randIndex := rand.Int31n(int32(len(keys)))
	randKey := keys[randIndex].String()
	randElement := resources.WebsiteData[randKey]

	return randElement
}

func main() {
	app.Route("/", &website{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "homepage tinder",
		Description: "Tinder swiping through a list of websites homepage's",
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
