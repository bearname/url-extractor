package app

import "sync"

type Crawler struct {
	Crawled map[string]bool
	mux     sync.Mutex
}


func New() *Crawler {
	return &Crawler{
		Crawled: make(map[string]bool),
	}
}

func (c *Crawler) visit(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	_, ok := c.Crawled[url]
	if ok {
		return true
	}
	c.Crawled[url] = true

	return false
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup

	v := c.visit(url)
	if v || depth <= 0 {
		return
	}

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		//fmt.Println(err)
		return
	}
	//fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			c.Crawl(u, depth-1, fetcher)
		}(u)
	}
	wg.Wait()
	return
}
