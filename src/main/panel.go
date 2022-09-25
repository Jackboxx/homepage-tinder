package main

import (
	"fmt"
	"math/rand"
	"reflect"

	"homepage-tinder/src/resources"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type WebsitePanel struct {
	app.Compo
	Name string
	Url  string
}

func (wp *WebsitePanel) Render() app.UI {
	// var element resources.Data
	// wp.Name, element = readRandomElement()
	// wp.Url = element.Url

	wp.Name = "ANDRITZ AG"
	wp.Url = "www.andritz.com"

	return app.Div().Body(
		app.Div().Body(
			app.Img().Src(fmt.Sprint("web/images/", wp.Url, ".png")),
		),
		app.Div().Body(
			app.Span().Text(wp.Name),
		),
	)
}

func readRandomElement() (string, resources.Data) {
	keys := reflect.ValueOf(resources.WebsiteData).MapKeys()
	randIndex := rand.Int31n(int32(len(keys)))
	randKey := keys[randIndex].String()
	randElement := resources.WebsiteData[randKey]

	return randKey, randElement
}
