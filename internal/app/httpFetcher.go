package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type HttpFetcher struct {
}

func (f *HttpFetcher) Fetch(urlLink string) (string, []string, error) {
	var (
		err       error
		links     = make([]string, 0)
		matches   [][]string
		findLinks = regexp.MustCompile("<a.*?href=\"(.*?)\"")
	)
	if !strings.Contains(urlLink, "https://") && !strings.Contains(urlLink, "http://") {
		return "", nil, err
	}
	fmt.Println(urlLink)
	resp, err := http.Get(urlLink)
	if err != nil {
		return "", nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	matches = findLinks.FindAllStringSubmatch(string(body), -1)

	for _, val := range matches {
		var linkUrl *url.URL

		if linkUrl, err = url.Parse(val[1]); err != nil {
			return string(body), links, err
		}

		if linkUrl.IsAbs() {
			links = append(links, linkUrl.String())
		} else {
			links = append(links, urlLink)
		}
	}

	return string(body), links, err
}
