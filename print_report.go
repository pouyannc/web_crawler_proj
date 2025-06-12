package main

import (
	"fmt"
	"slices"
)

type pagesStruct struct {
	url string
	amount int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("====================================================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("====================================================")

	sortedPages := sortPagesIntoStructSlice(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %v internal links to %s\n", page.amount, page.url)
	}
}

func sortPagesIntoStructSlice(pages map[string]int) []pagesStruct {
	sortedPages := []pagesStruct{}
	
	for k, v := range pages {
		sortedPages = append(sortedPages, pagesStruct{
			url: k,
			amount: v,
		})
	}

	slices.SortFunc(sortedPages, func(a, b pagesStruct) int {
		if a.amount != b.amount {
			return b.amount - a.amount
		}
		if a.url < b.url {
			return -1
		} else if a.url > b.url {
			return 1
		}
		return 0
	})

	return sortedPages
}
