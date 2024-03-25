package main

import (
	"github.com/vatusa/training/internal/web"
	"log"
)

func main() {
	err := web.Echo()
	if err != nil {
		log.Fatal(err)
		return
	}
}
