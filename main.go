// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/aprendagolang/webform/handlers"
// )

// func main() {
// 	http.HandlerFunc("/", handlers.SubscriptionHandler())

// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}

// }

// // um teste sobre o que fazer quando esta fazendo algo planejando um script para desenvolver commits com pequenas alterações

package main

import (
	"log"
	"net/http"

	"github.com/aprendagolang/webform/handlers"
)

func main() {
	http.HandlerFunc("/", handlers.SubscriptionHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

// um teste sobre o que fazer quando esta fazendo algo planejando um script para desenvolver commits com pequenas alterações
