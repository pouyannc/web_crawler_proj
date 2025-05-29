package main

import (
	"net/url"
	"strings"
)

func normalizeURL(URL string) (string, error) {
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	normHost := parsedURL.Host
	if normHost[:4] == "www." {
		normHost = normHost[4:]
	}
	normPath := strings.TrimRight(parsedURL.Path, "/")

	normalized := strings.TrimSpace(normHost + normPath)
	normalized = strings.ToLower(normalized)
	return normalized, nil
}