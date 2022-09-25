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
	Name        string
	Url         string
	openedPages []string
}

func (wp *WebsitePanel) Render() app.UI {
	if wp.Name == "" {
		var element resources.Data
		wp.Name, element = wp.readRandomElement()
		wp.Url = element.Url
	}

	return app.Div().Body(
		app.Div().Body(
			app.Span().Text(wp.Name),
		),
		app.Div().Body(
			app.Img().Src(fmt.Sprint("web/images/", wp.Url, ".png")),
		),
		app.Div().Body(
			app.Span().Text("Next"),
		).OnClick(wp.getNew),
		app.Div().Body(
			app.Span().Text("Open in new tab"),
		).OnClick(wp.onClick),
	)
}

func (wp *WebsitePanel) getNew(ctx app.Context, e app.Event) {
	oldName := wp.Name

	for {
		var element resources.Data
		wp.Name, element = wp.readRandomElement()
		wp.Url = element.Url

		if wp.Name != oldName {
			break
		}
	}
}

func (wp *WebsitePanel) onClick(ctx app.Context, e app.Event) {
	if wp.Url != "" {
		app.Window().Call("open", fmt.Sprint("https://", wp.Url))
	}
}

func (wp *WebsitePanel) readRandomElement() (string, resources.Data) {
	keys := reflect.ValueOf(resources.WebsiteData).MapKeys()
	var lastElement string
	if len(keys) == len(wp.openedPages) {
		lastElement = wp.openedPages[len(wp.openedPages)-1]
		wp.openedPages = []string{}
	}

	for {
		randIndex := rand.Int31n(int32(len(keys)))
		randKey := keys[randIndex].String()
		randElement := resources.WebsiteData[randKey]

		if !isElementInArray(wp.openedPages, randElement.Url) && randElement.Url != lastElement {
			wp.openedPages = append(wp.openedPages, randElement.Url)
			return randKey, randElement
		}
	}
}

func isElementInArray(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
