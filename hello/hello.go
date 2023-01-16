package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("hello")
	link, _ := html.Parse(strings.NewReader("hello"))
	fmt.Println(link)
}
