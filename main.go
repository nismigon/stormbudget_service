package main

import (
	"fmt"
	"log"

	"stormlab.fr/stormbudget/v1/server/configuration"
)

func main() {
	_, err := configuration.ParseConfiguration("configuration.json")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Pass")
	}
}
