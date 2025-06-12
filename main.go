package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages map[string]int
	baseURL *url.URL
	mu *sync.Mutex
	concurrencyControl chan struct{}
	wg *sync.WaitGroup
	maxPages int
}

var maxConcurrency = 5
var maxPages = 20

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawURL := args[0]
	var err error
	if len(args) == 2 {
		maxConcurrency, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("max concurrency argument provided is not a valid number")
			os.Exit(1)
		}
	}
	if len(args) == 3 {
		maxPages, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("max pages argument provided is not a valid number")
			os.Exit(1)
		}
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg := config {
		pages: map[string]int{},
		baseURL: u,
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg: &sync.WaitGroup{},
		maxPages: maxPages,
	}

	fmt.Printf("starting crawl of: %s\n", rawURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()

	for k, v := range cfg.pages {
		fmt.Println(k, v)
	}
}