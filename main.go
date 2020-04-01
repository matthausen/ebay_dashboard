package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers"
)

func main() {
	r := handlers.Router()

	fmt.Println("Starting server on the 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
