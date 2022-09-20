package main

import (
	"log"
	"wallet/api"
	"wallet/dep"
)

func main() {
	c, err := dep.Init()
	if err != nil {
		log.Fatal(err)
	}

	api.Run(c)
}
