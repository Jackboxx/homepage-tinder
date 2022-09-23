package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"

	"homepage-tinder/src/urlFetcher"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World!")
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, fmt.Sprintf("https://%s", url))
	return exec.Command(cmd, args...).Start()
}

// app.Route("/", &hello{})
// app.RunWhenOnBrowser()

// http.Handle("/", &app.Handler{
// 	Name:        "Hello",
// 	Description: "An Hello World! example",
// })
// if err := http.ListenAndServe(":8000", nil); err != nil {
// 	log.Fatal(err)
// }

func writeData(jsonData map[string]interface{}) {
	data, err := json.Marshal(jsonData)

	if err != nil {
		log.Print("Error during json Marshal: ", err)
	}

	err = ioutil.WriteFile("src/resources/data.json", data, 0644)

	if err != nil {
		log.Print("Error during file write: ", err)
	}
}

func readJson() {
	content, err := ioutil.ReadFile("src/resources/data.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json Unmarshal(): ", err)
	}

	for element := range payload {
		if payload[element] == nil {
			out, err := urlFetcher.BuildUrl(element)

			if err != nil {
				out = urlFetcher.QueryUrl(element)
				if out != "" {
					payload[element] = out
					fmt.Println(element + " : " + out)
					writeData(payload)
				}
			} else {
				payload[element] = out
				fmt.Println(element + " : " + out)
				writeData(payload)
			}
		}
	}

}

func main() {
	readJson()
}
