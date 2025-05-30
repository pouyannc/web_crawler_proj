package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	_, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	var refURLs []string


	htmlReader := strings.NewReader(htmlBody)
	htmlNodes, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	for n := range htmlNodes.Descendants() {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil {
						break
					}
					if u.IsAbs() {
						refURLs = append(refURLs, a.Val)
					} else {
						refURLs = append(refURLs, rawBaseURL + a.Val)
					}
					break
				}
			}
		}
	}

	return refURLs, nil
}