package main

import (
	"log"
	"net/http"

	"github.com/aprendagolang/webform/handlers"
)

func main() {
	http.HandlerFunc("/", handlers.SubscriptionHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
