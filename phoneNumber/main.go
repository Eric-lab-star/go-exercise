package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Print("hello")

}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")

}
