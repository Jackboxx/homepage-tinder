package main

import (
	"encoding/json"
	"fmt"
	"homepage-tinder/src/cli"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type JsonData struct {
	Description struct {
		Url string
	}
}

func getScreenshot() {

}

func writeData(jsonData map[string]JsonData) {
	data, err := json.Marshal(jsonData)

	if err != nil {
		log.Print("Error during json Marshal: ", err)
	}

	err = ioutil.WriteFile("src/resources/data.json", data, 0644)

	if err != nil {
		log.Print("Error during file write: ", err)
	}
}

func fillJson() {
	content, err := ioutil.ReadFile("src/resources/data.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload map[string]JsonData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during json Unmarshal(): ", err)
	}

	for element := range payload {
		url := payload[element].Description.Url
		if url == "" {
			out, err := cli.BuildUrl(element)

			if err != nil {
				if out != "" {
					copy := payload[element]
					copy.Description.Url = out
					fmt.Println(element, " ", out)
					payload[element] = copy
					writeData(payload)
				}
			} else {
				copy := payload[element]
				copy.Description.Url = out
				fmt.Println(element, " ", out)
				payload[element] = copy
				writeData(payload)
			}
		}

		_, err := os.Open(fmt.Sprint("web/images/", payload[element].Description.Url, ".png"))

		if err != nil && payload[element].Description.Url != "" {
			location := "src/cli/get-screenshot.js"
			url := fmt.Sprint("https://", payload[element].Description.Url)

			fmt.Println(url)
			exec.Command("node", location, url, payload[element].Description.Url).Run()
		}
	}
}

func main() {
	fillJson()
}
