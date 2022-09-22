package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Escaper struct {
	escaped   string
	unescaped string
}

func getUrl(query string) (string, error) {
	q := replaceSpaces(query)
	args := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s&s=&dc=&v=l&o=json", q)
	data, err := exec.Command("curl", args).Output()

	if err != nil {
		return "", err
	}

	element := trimRemainingLeft(string(data), "result__a")
	first_result := trimRemainingRight(element, "</a>")
	url_escaped := parseEscapedUrl(first_result)
	url := unescapeString(url_escaped)

	return url, nil
}

func parseEscapedUrl(str string) string {
	first_step := trimRemainingLeft(str, "uddg=")
	second_step := trimRemainingRight(first_step, "\"")
	third_step := trimRemainingRight(second_step, "&")

	return third_step
}

func unescapeString(str string) string {
	sequences := [25]Escaper{}

	sequences[0].escaped = "%2F"
	sequences[0].unescaped = "/"

	sequences[1].escaped = "%3A"
	sequences[1].unescaped = ":"

	sequences[2].escaped = "%2D"
	sequences[2].unescaped = "-"

	sequences[3].escaped = "%20"
	sequences[3].unescaped = " "

	sequences[4].escaped = "%24"
	sequences[4].unescaped = "$"

	sequences[4].escaped = "%26"
	sequences[4].unescaped = "&"

	sequences[5].escaped = "%60"
	sequences[5].unescaped = "`"

	sequences[6].escaped = "%3C"
	sequences[6].unescaped = "<"

	sequences[7].escaped = "%3E"
	sequences[7].unescaped = ">"

	sequences[8].escaped = "%5B"
	sequences[8].unescaped = "["

	sequences[9].escaped = "%5D"
	sequences[9].unescaped = "]"

	sequences[10].escaped = "%7B"
	sequences[10].unescaped = "{"

	sequences[11].escaped = "%7D"
	sequences[11].unescaped = "}"

	sequences[12].escaped = "%22"
	sequences[12].unescaped = "\""

	sequences[13].escaped = "%23"
	sequences[13].unescaped = "#"

	sequences[14].escaped = "%25"
	sequences[14].unescaped = "%"

	sequences[15].escaped = "%40"
	sequences[15].unescaped = "@"

	sequences[16].escaped = "%3B"
	sequences[16].unescaped = ";"

	sequences[17].escaped = "%3D"
	sequences[17].unescaped = "="

	sequences[18].escaped = "%3F"
	sequences[18].unescaped = "?"

	sequences[19].escaped = "%5C"
	sequences[19].unescaped = "\\"

	sequences[20].escaped = "%5E"
	sequences[20].unescaped = "^"

	sequences[21].escaped = "%7C"
	sequences[21].unescaped = "|"

	sequences[22].escaped = "%7E"
	sequences[22].unescaped = "~"

	sequences[23].escaped = "%27"
	sequences[23].unescaped = "'"

	sequences[24].escaped = "%2C"
	sequences[24].unescaped = ","

	res := str

	for i, _ := range sequences {
		res = strings.ReplaceAll(res, sequences[i].escaped, sequences[i].unescaped)
	}

	return res
}

func replaceSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "+")
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
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func intMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trimRemainingRight(str string, cutoff string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		search := ""
		limit := intMin(i+len(cutoff), len(str))

		for j := i; j < limit; j++ {
			if j < 0 {
				break
			}
			search += string(str[j])
			if search == cutoff {
				return result
			}
		}

		result += string(str[i])
	}

	return result
}

func trimRemainingLeft(str string, cutoff string) string {
	result := ""
	reached := false

	for i := 0; i < len(str); i++ {

		if !reached {
			search := ""
			limit := i - len(cutoff)
			for j := limit; j < i; j++ {
				if j < 0 {
					break
				}
				search += string(str[j])
				if search == cutoff {
					reached = true
				}
			}
		}

		if reached {
			result += string(str[i])
		}
	}

	return result
}

func shuffleSlice(slice []string) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func returnUrl() {
	var queries []string
	var filename string

	fmt.Print("Enter file location: ")
	fmt.Scan(&filename)

	rfile, frErr := os.Open(filename)

	if frErr != nil {
		log.Fatal("Failed to open file")
	}

	wfile, fwErr := os.Create("urls")

	if fwErr != nil {
		log.Fatal("Failed to create file")
	}

	scanner := bufio.NewScanner(rfile)
	for scanner.Scan() {
		queries = append(queries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Failed to read file")
	}

	shuffleSlice(queries)

	var url string = ""
	var err error

	for _, s := range queries {
		timeout := 1.0

		for {
			channel := make(chan string)

			go func() {
				url, err = getUrl(s)

				if err != nil {
					log.Fatal("Failed to get Url")
				}

				channel <- url
			}()

			result_url := <-channel

			if len(result_url) > 0 {
				break
			}

			time.Sleep(time.Duration(timeout) * time.Second)
			timeout = math.Min(timeout*1.5, 10)
		}

		fmt.Printf("wrote URL: %s to URL file\n", url)
		wfile.Write([]byte(fmt.Sprintf("%s\n", url)))
	}

	defer rfile.Close()
	defer wfile.Close()
	fmt.Println("Finished")
}
