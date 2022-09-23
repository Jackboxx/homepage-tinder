package cli

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

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
	url := url.PathEscape(url_escaped)

	return url, nil
}

func parseEscapedUrl(str string) string {
	first_step := trimRemainingLeft(str, "uddg=")
	second_step := trimRemainingRight(first_step, "\"")
	third_step := trimRemainingRight(second_step, "&")

	return third_step
}

func replaceSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "+")
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

func QueryUrl(name string) string {
	var url string
	var err error

	for {
		channel := make(chan string)

		go func() {
			url, err = getUrl(name)

			if err != nil {
				log.Print(err)
			} else {
				channel <- url
			}
		}()

		result_url := <-channel

		if len(result_url) > 0 {
			break
		}

		time.Sleep(20 * time.Second)
		return ""
	}

	return trimRemainingLeft(trimRemainingLeft(url, "//"), "www.")
}
