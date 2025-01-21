package main

import (
	"log"
	"strings"
)

func main() {
	urls := strings.Split("/api/platform/create", "/")

	log.Println(strings.Join(urls[2:], "/"))
}
