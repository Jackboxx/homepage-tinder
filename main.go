package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

type Escaper struct {
    escaped string
    unescaped string
}


func getUrl(query string) (string, error) {
    q := replaceSpaces(query)
    args := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s&s=&dc=&v=l&o=json", q)
    data, err := exec.Command("curl", args).Output()

    if err != nil {
        return "", err;
    }

    element := strings.TrimLeft(string(data), "result__a")
    first_result := strings.TrimLeft(element, "</a>")
    url_escaped := parseEscapedUrl(first_result)
    url := unescapeString(url_escaped)

    return url, nil
}

func parseEscapedUrl(str string) string {
    first_step := strings.TrimLeft(str, "uddg=")
    // fmt.Println(first_step)
    second_step := strings.TrimRight(first_step,  "\"")
    // fmt.Println(second_step)
    third_step := strings.TrimRight(second_step, "&")
    // fmt.Println(third_step)

    return third_step
}

func unescapeString(str string) string {
    sequences := [3]Escaper{}

    sequences[0].escaped = "%2F"
    sequences[0].unescaped = "/"

    sequences[1].escaped = "%3A"
    sequences[1].unescaped = ":"

    sequences[2].escaped = "%2D"
    sequences[2].unescaped = "-"

    res := str

    for i, _ := range sequences{
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

func trimRemainingRight(str string, cutoff string) string {
    result := ""
    for i := 0; i < len(str); i++ {
        search := ""
        limit := i - len(cutoff)
        for j := limit; j < i; j++ {
            if(j < 0) { break }
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
        if (reached) {
            result += string(str[i])
        }

        if (!reached) {
            search := ""
            limit := i - len(cutoff) 
            for j := limit; j < i; j++ {
                fmt.Println(search)
                if(j < 0) { break } 
                search += string(str[j])
                if search == cutoff {
                    reached = true
                }
            }
        }
    }

    return result
}

func main() {

    test := "test more"
    fmt.Print("re: " + trimRemainingLeft(test, "mo"))

    url, err := getUrl("a test query")

    if(err != nil) {
        log.Fatal("Failed to get Url")
    }

    open(url)
}
