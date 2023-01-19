package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

	return filterAll(hrefs(res.Body, baseURL), withNumber(100), withNumber(90))
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

func withNumber(number int) func(string) bool {
	return func(href string) bool {
		return strings.Contains(href, strconv.Itoa(number))

	}
}
