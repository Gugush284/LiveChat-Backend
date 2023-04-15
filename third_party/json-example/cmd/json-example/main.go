package main

import (
	"LiveChat-Backend/third_party/json-example/internal/jsonServer"
	"log"
)

func main() {

	config := jsonServer.NewConfig()

	if err := jsonServer.Start(config); err != nil {
		log.Fatal(err)
	}
}
