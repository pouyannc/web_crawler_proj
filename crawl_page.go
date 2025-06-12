package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

	cfg.concurrencyControl <- struct{}{}
	defer func (){<-cfg.concurrencyControl}()

	if cfg.maxPagesReached() {
		return
	}

	currentU, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing raw current URL: %v", err)
		return
	}
	if cfg.baseURL.Hostname() != currentU.Hostname(){
		return
	}

	currentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing current URL: %v", err)
		return
	}
	
	if !cfg.addPageVisit(currentURL) {
		return
	}

	fmt.Printf("Crawling new page: %s\n", currentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting HTML from current URL: %v", err)
		return
	}
	urlsFromHTML, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("couldn't get urls from current HTML: %v", err)
		return
	}
	for _, urlFromHTML := range urlsFromHTML {
		cfg.wg.Add(1)
		go cfg.crawlPage(urlFromHTML)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) maxPagesReached() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) >= cfg.maxPages
}