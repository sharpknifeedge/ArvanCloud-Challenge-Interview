package main

import (
	"log"
	"voucher/api"
	"voucher/dep"
)

func main() {
	c, err := dep.Init()
	if err != nil {
		log.Fatal(err)
	}
	api.Run(c)
}
