package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Gugush284/"
)

func main() {
	flag.Parse()

	config := apiserver.NewConfig() 

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}