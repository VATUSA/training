package main

import (
	"log"
	"vatusa-training/internal/web"
)

func main() {
	err := web.Echo()
	if err != nil {
		log.Fatal(err)
		return
	}
}
