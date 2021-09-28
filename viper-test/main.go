package main

import (
	"fmt"
	"log"
	"viper-test/config"
)

func main() {
	configFile, err := config.Load(".", "env")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	fmt.Printf("%+v\n", configFile)
}
