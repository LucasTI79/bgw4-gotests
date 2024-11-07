package main

import (
	"log"
	"net/http"

	"github.com/lucasti79/bgw4-put-patch-delete/config"
	"github.com/lucasti79/bgw4-put-patch-delete/internal/server"
)

// TDD -> Teste depois do deploy

func main() {
	config.Init()
	db := config.InitDatabase()
	server := server.CreateServer(db)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
