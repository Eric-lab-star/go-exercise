package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Eric-lab-star/go-exercise/link"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url of the website you want to build")
	flag.Parse()

	hrefs := get(*urlFlag)

	for _, href := range hrefs {
		fmt.Println(href)
	}
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

	return filter(baseURL.String(), hrefs(res.Body, baseURL))
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

func filter(base string, hrefs []string) []string {
	var ret []string
	for _, href := range hrefs {
		switch {
		case strings.HasPrefix(href, base):
			ret = append(ret, href)
		}
	}
	return ret
}
