package main

import (
	"app/internal/server"
	"embed"
	"log"
)

var embeddedFrontend embed.FS

func main() {
	if err := server.Start(":8080", embeddedFrontend); err != nil {
		log.Println("server stopped:", err)
	}
}
