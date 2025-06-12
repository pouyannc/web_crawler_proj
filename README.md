# Web Crawler

Crawls a given webpage URL, outputting a report to the CLI of all internal links and the number of times they were found.

## Getting Started

Requires Go 1.21+
Clone the repository and build the Go binary:

```bash
git clone https://github.com/pouyannc/web_crawler_proj.git
cd web_crawler_proj
go build -o crawler
```

## Usage

After building the binary, run the program with the URL of the site you want to begin crawling as the first argument:

```bash
./crawler <website URL>
```

By default, the max concurrency of go routines is set to 5, and the max pages it will crawl is 20.
These can be changed by providing the custom values as arguments:

```bash
./crawler <website URL> <concurrency limit> <page limit>
```
