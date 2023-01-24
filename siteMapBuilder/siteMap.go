package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Eric-lab-star/go-exercise/link"
)

type loc struct {
	Value string `xml:"loc"`
}

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url of the website you want to build")
	maxDepth := flag.Int("depth", 5, "the maximum number of links deep to traverse")
	flag.Parse()
	links := bfs(*urlFlag, *maxDepth)
	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, link := range links {
		toXml.Urls = append(toXml.Urls, loc{link})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
	fmt.Println()

}

func bfs(url string, maxdeep int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		url: {},
	}
	for i := 0; i < maxdeep; i++ {
		q, nq = nq, make(map[string]struct{})
		for l := range q {
			if _, ok := seen[l]; ok {
				continue
			}
			seen[l] = struct{}{}
			for _, link := range get(l) {
				nq[link] = struct{}{}
			}
		}
	}
	var ret []string
	for link := range seen {
		ret = append(ret, link)
	}
	return ret

}

// get makes http.Get request and return slice of href string
func get(urlFlag string) []string {
	res, err := http.Get(urlFlag)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	reqURL := res.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	return filterAll(hrefs(res.Body, baseURL), withPrefix(baseURL.String()))
}

// hrefs parse response body and return href
func hrefs(r io.Reader, baseURL *url.URL) []string {

	links, _ := link.Parse(r)
	var ret []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			ret = append(ret, baseURL.String()+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			ret = append(ret, link.Href)
		}
	}

	return ret
}

func filter(hrefs []string, keepFn func(string) bool) []string {
	var ret []string

	for _, href := range hrefs {
		if keepFn(href) {
			ret = append(ret, href)
		}

	}
	return ret
}

func filterAll(hrefs []string, keepFns ...func(string) bool) []string {
	var ret []string

	for i := 0; i < len(keepFns); i++ {
		if i == 0 {
			ret = filter(hrefs, keepFns[i])
			continue
		}
		ret = filter(ret, keepFns[i])
	}

	return ret
}

func withPrefix(str string) func(string) bool {

	return func(href string) bool {
		if strings.HasPrefix(href, str) {
			return true
		} else {
			return false
		}
	}

}
