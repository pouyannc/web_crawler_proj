package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int, error) {
	baseU, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing raw base URL: %v", err)
		return pages, err
	}
	currentU, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing raw current URL: %v", err)
		return pages, err
	}
	if baseU.Hostname() != currentU.Hostname(){
		return pages, nil
	}

	currentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing current URL: %v", err)
		return pages, err
	}
	
	if _, ok := pages[currentURL]; ok {
		pages[currentURL]++
		return pages, nil
	}

	pages[currentURL] = 1
	fmt.Printf("Crawling new page: %s\n", currentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting HTML from current URL: %v", err)
		return pages, nil
	}
	urlsFromHTML, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("couldn't get urls from current HTML: %v", err)
		return pages, nil
	}
	for _, urlFromHTML := range urlsFromHTML {
		pages, err = crawlPage(rawBaseURL, urlFromHTML, pages)
		if err != nil {
			return pages, err
		}
	}

	return pages, nil
}