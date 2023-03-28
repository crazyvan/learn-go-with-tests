package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":5000", server))
}
