package main

import (
	"flag"
	"fmt"
	"github.com/bearname/url-extractor/pkg/app"
	"github.com/bearname/url-extractor/pkg/infrastructure/util"
)

func main() {
	url := flag.String("url", "localhost", "a string")
	depth := flag.Int("depth", 4, "depth of crawl an int, default 4")
	crawler := app.New()
	elapsed := util.Elapsed(func() {
		crawler.Crawl(*url, *depth, &app.HttpFetcher{})
	})

	for key := range crawler.Crawled {
		fmt.Println(key)
	}
	fmt.Println(len(crawler.Crawled))

	fmt.Printf("Elapsed time %d", elapsed)
}
