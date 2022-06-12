package main

import (
	"fmt"
	"net/http"

	"github.com/webhook-repo/handlers"
)

func main() {

	http.HandleFunc("/webhook", handlers.WebhookHandler)
	fmt.Println("Start...")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
