package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("get request for webpage returned 400+ status code")
	}

	if !strings.Contains(res.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("get request for webpage returned non-html content type")
	} 

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}