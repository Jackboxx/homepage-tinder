package cli

import (
	"errors"
	"fmt"
	"strings"
)

func pingDomain(domain string) bool {
	return false
}

func build(name string, prefix string, suffix string) (string, error) {
	domain := fmt.Sprintf("%s%s%s", prefix, name, suffix)

	if pingDomain(domain) {
		return domain, nil
	}

	return "", errors.New("No response")
}

func BuildUrl(description string) (string, error) {
	name := (strings.Split(strings.ToLower(strings.Trim(description, " ")), " ")[0])
	var domain string
	var err error

	domain, err = build(name, "www.", ".at")
	if err == nil {
		return domain, err
	}

	domain, err = build(name, "www.", ".com")
	if err == nil {
		return domain, err
	}

	domain, err = build(name, "", ".at")
	if err == nil {
		return domain, err
	}

	domain, err = build(name, "", ".com")
	if err == nil {
		return domain, err
	}

	words := strings.Split(strings.ToLower(strings.Trim(description, " ")), " ")
	if len(words) >= 2 {
		name = fmt.Sprintf("%s-%s", words[0], words[1])

		domain, err = build(name, "www.", ".at")
		if err == nil {
			return domain, err
		}

		domain, err = build(name, "www.", ".com")
		if err == nil {
			return domain, err
		}

		domain, err = build(name, "", ".at")
		if err == nil {
			return domain, err
		}

		domain, err = build(name, "", ".com")
		if err == nil {
			return domain, err
		}
	}

	errString := fmt.Sprintf("Failed to build url for %s\n", description)
	return "", errors.New(errString)
}
