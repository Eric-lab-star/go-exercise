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

	hrefs := hrefHandler(*urlFlag)

	for _, href := range hrefs {
		fmt.Println(href)
	}
}

func hrefHandler(urlFlag string) []string {
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

	return hrefs(res.Body, baseURL)
}

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
