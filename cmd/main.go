package main

import (
	"crawler/internal/app"
	"crawler/internal/infrastructure/util"
	"flag"
	"fmt"
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
