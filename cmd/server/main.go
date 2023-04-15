package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()

	config := server.NewConfig()

	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}
