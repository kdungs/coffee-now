package main

import (
	"log"
	"net/http"
)

func main() {
	cns, err := NewCoffeeNowServer()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", cns))
}
